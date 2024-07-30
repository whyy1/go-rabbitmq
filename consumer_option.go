package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq-pool/internal"
	"go.uber.org/zap"
	"time"
)

type ConsumerOptions struct {
	ReconnectInterval time.Duration
	Logger            internal.Logger
	ExchangeOptions   ExchangeOptions
	QueueOptions      QueueOptions
	ConsumOptions     ConsumOptions
}
type QueueOptions struct {
	Name       string //队列名称
	Durable    bool   //队列是否持久化
	AutoDelete bool   //是否自动删除
	Exclusive  bool   //是否具有排他性
	NoWait     bool   //是否阻塞处理
	Passive    bool
	Args       amqp.Table //额外的属性
	Declare    bool
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
