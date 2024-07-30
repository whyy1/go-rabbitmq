package main

import (
	"errors"
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

	if err := declareExchange(chanManager, options.ExchangeOptions); err != nil {
		return nil, err
	}

	consumer = &Consumer{
		connectionManager: connectionManager,
		chanManager:       chanManager,
		options:           options,
	}

	return
}

func (consumer *Consumer) Close() {
	consumer.chanManager.CloseChannel()
}
