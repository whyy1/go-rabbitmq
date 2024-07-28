package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq-pool/internal"
	"time"
)

// ExchangeOptions are used to configure an exchange.
// If the Passive flag is set the client will only check if the exchange exists on the server
// and that the settings match, no creation attempt will be made.
type ExchangeOptions struct {
	Name       string
	Kind       string // possible values: empty string for default exchange or direct, topic, fanout
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Passive    bool // if false, a missing exchange will be created on the server
	Args       amqp.Table
	Declare    bool
	//Bindings   []amqp.Binding
}
type PublisherOptions struct {
	ReconnectInterval time.Duration
	Logger            internal.Logger
	ExchangeOptions   ExchangeOptions
}

func WithDefaultPublishOptionsOptions(connectionOptions *internal.ConnectionOptions) (options PublisherOptions) {
	return PublisherOptions{
		ReconnectInterval: connectionOptions.ReconnectInterval,
		Logger:            connectionOptions.Logger,
		ExchangeOptions: ExchangeOptions{
			Name:       "",
			Kind:       amqp.ExchangeDirect,
			Durable:    false,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Passive:    false,
			Args:       amqp.Table{},
			Declare:    false,
		},
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
