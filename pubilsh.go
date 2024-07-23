package main

import (
	"go_rabbitmq_pool/internal"
)

type Publisher struct {
	connectionManager *internal.ConnectionManager
	chanManager       *internal.ChannelManager
	options           internal.ConnectionOptions
}

func NewPubilsh(connectionManager *internal.ConnectionManager) (publisher *Publisher, err error) {

	manager, err := internal.NewChannelManager(connectionManager)
	if err != nil {
		return nil, err
	}
	publisher = &Publisher{
		connectionManager: connectionManager,
		chanManager:       manager,
		options:           internal.ConnectionOptions{},
	}

	return
}

func (publisher *Publisher) PublishWithContext(exchange string, key string, mandatory bool, immediate bool, data []byte) {

	publisher.chanManager.PublishWithContext(exchange, key, mandatory, immediate, data)
}

func (publisher *Publisher) Close() {
	publisher.chanManager.ChannelClose()
}
