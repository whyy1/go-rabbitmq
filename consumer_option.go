package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/whyy1/go-rabbitmq/internal"
	"go.uber.org/zap"
	"time"
)

type ConsumerOptions struct {
	ReconnectInterval time.Duration
	Logger            internal.Logger
	ExchangeOptions   ExchangeOptions
	QueueOptions      QueueOptions
	ConsumeOptions    ConsumeOptions
	QueueBindOptions  QueueBindOptions
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

type QueueBindOptions struct {
	Name     string     //队列名称， 要绑定的队列名称。
	Key      string     // 指定消息的路由键。交换机会根据这个键来决定将消息发送到哪个队列。
	Exchange string     //指定将要绑定的交换机名称。消息将从这个交换机路由到绑定的队列。
	NoWait   bool       //如果设置为 true，表示不等待服务器的响应。绑定操作立即返回，不会确认绑定是否成功。
	Args     amqp.Table //传递一些额外的绑定参数，可以用于配置更复杂的绑定规则或扩展功能。例如，可以设置绑定的优先级等
	Bind     bool       //是否绑定至交换机
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
func WithConsumerBindKey(key string) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueBindOptions.Key = key
	}
}
func WithConsumerBindNoWait(noWait bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueBindOptions.NoWait = noWait
	}
}
func WithConsumerBindArgs(args amqp.Table) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueBindOptions.Args = args
	}
}
func WithConsumerBind(bind bool) func(options *ConsumerOptions) {
	return func(options *ConsumerOptions) {
		options.QueueBindOptions.Bind = bind
	}
}
