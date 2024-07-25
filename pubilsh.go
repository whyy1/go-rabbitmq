package main

import (
	"github.com/whyy1/go-rabbitmq-pool/internal"
	"go.uber.org/zap"
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

	if err := publisher.chanManager.PublishWithContext(exchange, key, mandatory, immediate, data); err != nil {
		publisher.options.Logger.Errorf("消息发送失败，", zap.Error(err))
		return
	}
}

func (publisher *Publisher) Close() {
	publisher.chanManager.CloseChannel()
}
