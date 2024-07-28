package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq-pool/internal"
	"go.uber.org/zap"
)

type Publisher struct {
	connectionManager *internal.ConnectionManager
	chanManager       *internal.ChannelManager
	options           PublisherOptions
}

func NewPubilsh(connectionManager *internal.ConnectionManager) (publisher *Publisher, err error) {
	if connectionManager == nil {
		return nil, errors.New("connectionManager is nil")
	}
	defaultOptions := connectionManager.GetConnectionOptions()

	chanManager, err := internal.NewChannelManager(connectionManager)
	if err != nil {
		return nil, err
	}
	publisher = &Publisher{
		connectionManager: connectionManager,
		chanManager:       chanManager,
		options:           WithDefaultPublishOptionsOptions(&defaultOptions),
	}

	return
}

func (publisher *Publisher) PublishWithContext(
	ctx context.Context,
	data []byte,
	routingKeys []string,
	optionFuncs ...func(*PublishOptions),
) error {

	options := &PublishOptions{}
	for _, optionFunc := range optionFuncs {
		optionFunc(options)
	}
	message := amqp091.Publishing{
		Headers:         options.Headers,
		ContentType:     options.ContentType,
		ContentEncoding: options.ContentEncoding,
		DeliveryMode:    options.DeliveryMode,
		Priority:        options.Priority,
		CorrelationId:   options.CorrelationID,
		ReplyTo:         options.ReplyTo,
		Expiration:      options.Expiration,
		MessageId:       options.MessageID,
		Timestamp:       options.Timestamp,
		Type:            options.Type,
		UserId:          options.UserID,
		AppId:           options.AppID,
		//Body:            nil,
	}

	for _, routingKey := range routingKeys {
		//if err := publisher.chanManager.QueueDeclareSafe(routingKey, false, false, false, false, nil); err != nil {
		//	return err
		//}
		message.Body = data

		if err := publisher.chanManager.PublishWithContextSafe(
			ctx,
			options.Exchange,
			routingKey,
			options.Mandatory,
			options.Immediate,
			message); err != nil {
			publisher.options.Logger.Errorf("消息发送失败", zap.Error(err))
			return err
		}
	}
	fmt.Println("消息发送成功")
	publisher.options.Logger.Errorf("消息发送成功")
	return nil
}

func (publisher *Publisher) Close() {
	publisher.chanManager.CloseChannel()
}
