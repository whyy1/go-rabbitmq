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
	ConsumeOptions    ConsumeOptions
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
		ExchangeOptions:   getDefaultExchangeOptions(),
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

func WithConsumerExchangeName(name string) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Name = name
	}
}
func WithConsumerExchangeKind(kind string) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Kind = kind
	}
}
func WithConsumerExchangeDurable(durable bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Durable = durable
	}
}
func WithConsumerExchangeAutoDelete(autoDelete bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.AutoDelete = autoDelete
	}
}
func WithConsumerExchangeInternal(internal bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Internal = internal
	}
}
func WithConsumerExchangeNoWait(noWait bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.NoWait = noWait
	}
}
func WithConsumerExchangePassive(passive bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Passive = passive
	}
}

func WithConsumerExchangeArgs(args amqp.Table) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Args = args
	}
}
func WithConsumerExchangeDeclare(declare bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.ExchangeOptions.Declare = declare
	}
}

func WithConsumerQueueName(name string) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.Name = name
	}
}

func WithConsumerQueueDurable(durable bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.Durable = durable
	}
}

func WithConsumerQueueAutoDelete(autoDelete bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.AutoDelete = autoDelete
	}
}

func WithConsumerQueueExclusive(exclusive bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.Exclusive = exclusive
	}
}

func WithConsumerQueueNoWait(noWait bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.NoWait = noWait
	}
}
func WithConsumerQueuePassive(passive bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.Passive = passive
	}
}

func WithConsumerQueueArgs(args amqp.Table) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.Args = args
	}
}
func WithConsumerQueueDeclare(declare bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueOptions.Declare = declare
	}
}
