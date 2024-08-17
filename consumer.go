package main

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq/internal"
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
	optionFuncs ...func(*ConsumeOptions),
) (<-chan amqp.Delivery, error) {

	for _, optionFunc := range optionFuncs {
		optionFunc(&consumer.options.ConsumeOptions)
	}

	if err := declareExchange(consumer.chanManager, consumer.options.ExchangeOptions); err != nil {
		return nil, err
	}

	if err := declareQuene(consumer.chanManager, &consumer.options.QueueOptions); err != nil {
		return nil, err
	}
	consumer.options.QueueBindOptions.Exchange = consumer.options.ExchangeOptions.Name
	consumer.options.QueueBindOptions.Name = consumer.options.QueueOptions.Name

	if err := queneBind(consumer.chanManager, consumer.options.QueueBindOptions); err != nil {
		return nil, err
	}

	return consumer.chanManager.ConsumeWithContextSafe(
		ctx,
		consumer.options.QueueOptions.Name,
		consumer.options.ConsumeOptions.Name,
		consumer.options.ConsumeOptions.AutoAck,
		consumer.options.ConsumeOptions.Exclusive,
		consumer.options.ConsumeOptions.NoLocal,
		consumer.options.ConsumeOptions.NoWait,
		consumer.options.ConsumeOptions.Args,
	)
}

func (consumer *Consumer) Close() {
	consumer.chanManager.CloseChannel()
}
