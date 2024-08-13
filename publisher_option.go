package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq/internal"
	"go.uber.org/zap"
	"time"
)

type PublisherOptions struct {
	ReconnectInterval time.Duration
	Logger            internal.Logger
	ExchangeOptions   ExchangeOptions
}

func WithDefaultPublishOptionsOptions() (options PublisherOptions) {
	logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any

	return PublisherOptions{
		ReconnectInterval: 5 * time.Second,
		Logger:            logger.Sugar(),
		ExchangeOptions:   getDefaultExchangeOptions(),
	}
}

func WithReconnectInterval(interval time.Duration) func(options *PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ReconnectInterval = interval
	}
}
func WithLogger(logger internal.Logger) func(options *PublisherOptions) {
	return func(options *PublisherOptions) {
		options.Logger = logger
	}
}

// WithPublisherOptionsExchangeName sets the exchange name
func WithPublisherOptionsExchangeName(name string) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Name = name
	}
}

// WithPublisherOptionsExchangeKind ensures the queue is a durable queue
func WithPublisherOptionsExchangeKind(kind string) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Kind = kind
	}
}

// WithPublisherOptionsExchangeDurable ensures the exchange is a durable exchange
func WithPublisherOptionsExchangeDurable(durable bool) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Durable = durable
	}
}

// WithPublisherOptionsExchangeAutoDelete ensures the exchange is an auto-delete exchange
func WithPublisherOptionsExchangeAutoDelete(autoDelete bool) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.AutoDelete = autoDelete
	}
}

// WithPublisherOptionsExchangeInternal ensures the exchange is an internal exchange
func WithPublisherOptionsExchangeInternal(Internal bool) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Internal = Internal
	}
}

// WithPublisherOptionsExchangeNoWait ensures the exchange is a no-wait exchange
func WithPublisherOptionsExchangeNoWait(noWait bool) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.NoWait = noWait
	}
}

// WithPublisherOptionsExchangeDeclare will create the exchange if it doesn't exist
func WithPublisherOptionsExchangeDeclare(declare bool) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Declare = declare
	}
}

// WithPublisherOptionsExchangePassive ensures the exchange is a passive exchange
func WithPublisherOptionsExchangePassive(passive bool) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Passive = passive
	}
}

// WithPublisherOptionsExchangeArgs adds optional args to the exchange
func WithPublisherOptionsExchangeArgs(args amqp.Table) func(*PublisherOptions) {
	return func(options *PublisherOptions) {
		options.ExchangeOptions.Args = args
	}
}
