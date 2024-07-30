package main

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq-pool/internal"
)

type Consumer struct {
	connectionManager *internal.ConnectionManager
	chanManager       *internal.ChannelManager
	options           ConsumerOptions
}

func NewConsumer(connectionManager *internal.ConnectionManager, optionFuncs ...func(*ConsumerOptions)) (consumer *Consumer, err error) {
	if connectionManager == nil {
		return nil, errors.New("connectionManager is nil")
	}

	options := WithDefaultConsumerOptions()
	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}
	chanManager, err := internal.NewChannelManager(connectionManager)
	if err != nil {
		return nil, err
	}

	consumer = &Consumer{
		connectionManager: connectionManager,
		chanManager:       chanManager,
		options:           options,
	}

	return
}

func (consumer *Consumer) GetConsumeChannel(
	ctx context.Context,
) (<-chan amqp.Delivery, error) {

	if err := declareQuene(consumer.chanManager, consumer.options.QueueOptions); err != nil {
		return nil, err
	}

	if err := declareExchange(consumer.chanManager, consumer.options.ExchangeOptions); err != nil {
		return nil, err
	}

	return consumer.chanManager.ConsumeWithContextSafe(
		ctx,
		consumer.options.QueueOptions.Name,
		consumer.options.ConsumOptions.Name,
		consumer.options.ConsumOptions.AutoAck,
		consumer.options.ConsumOptions.Exclusive,
		consumer.options.ConsumOptions.NoLocal,
		consumer.options.ConsumOptions.NoWait,
		consumer.options.ConsumOptions.Args,
	)
}

func (consumer *Consumer) Close() {
	consumer.chanManager.CloseChannel()
}
