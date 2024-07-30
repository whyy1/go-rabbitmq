package main

import (
	"github.com/whyy1/go-rabbitmq-pool/internal"
	"go.uber.org/zap"
	"time"
)

type ConsumerOptions struct {
	ReconnectInterval time.Duration
	Logger            internal.Logger
	ExchangeOptions   ExchangeOptions
}

func WithDefaultConsumerOptions() (options ConsumerOptions) {
	logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any

	return ConsumerOptions{
		ReconnectInterval: 5 * time.Second,
		Logger:            logger.Sugar(),
		ExchangeOptions:   getExchangeOptions(),
	}
}

func WithConsumerReconnectInterval(interval time.Duration) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ReconnectInterval = interval
	}
}
func WithConsumerReconnectLogger(logger internal.Logger) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.Logger = logger
	}
}
