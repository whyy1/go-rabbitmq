package main

import amqp "github.com/rabbitmq/amqp091-go"

// ExchangeOptions 包含交换机配置选项
type ExchangeOptions struct {
	Name string // 名称
	Kind string // 类型：默认交换机留空或使用 "direct"，其他可能的值有 "topic"、"fanout"
	//(amq.direct) Direct Exchange 将消息路由到绑定键完全匹配的队列。这种交换机最常用于需要精确匹配路由键的场景。
	//(amq.fanout) Fanout Exchange 将消息路由到所有绑定到该交换机的队列，而忽略路由键。它类似于广播，将消息传递给所有订阅者。
	//(amq.match、amq.headers)  Headers Exchange 根据消息头属性而不是路由键来路由消息。可以匹配消息头的特定键值对。
	//("" or amq.default) Default Exchange 是一个隐式的直连交换机，每个队列自动绑定到它，绑定键与队列名称相同。

	Durable    bool       // 持久化：如果为 true，交换机在服务器重启后仍然存在
	AutoDelete bool       // 自动删除：如果为 true，当没有绑定队列时，交换机会自动删除
	Internal   bool       // 内部：如果为 true，交换机只能由其他交换机发布消息
	NoWait     bool       // 无需等待：如果为 true，声明交换机时不会等待服务器的响应
	Passive    bool       // 被动：如果为 false，服务器上缺少交换机时将自动创建，默认为false
	Args       amqp.Table // 参数：其他交换机的可选参数，通常用于扩展功能
	Declare    bool       // 声明：如果为 true，则声明此交换机
	Bindings   []string   // 绑定：交换机的绑定键列表
}

func getExchangeOptions() ExchangeOptions {
	return ExchangeOptions{
		Name:       "",
		Kind:       amqp.ExchangeDirect,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Passive:    false,
		Args:       amqp.Table{},
		Declare:    false,
	}
}
