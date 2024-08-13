package internal

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ExchangeDeclareSafe 用于声明一个交换机。如果交换机不存在，会创建一个新的交换机
// 如果交换机不存在，服务器会创建一个新的交换机。如果交换机已经存在且参数与现有交换机的参数不同，会返回一个错误。
func (channelManager *ChannelManager) ExchangeDeclareSafe(
	name string, // 交换机名称
	kind string, // 交换机类型
	durable bool, // 是否持久化
	autoDelete bool, // 是否自动删除
	internal bool, // 是否为内部交换机
	noWait bool, // 是否阻塞处理
	args amqp.Table, // 额外的属性
) error {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	return channelManager.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

// ExchangeDeclarePassiveSafe 用于检查交换机是否存会返回一个错误
// 不会创建新的交换机。仅用于检查交换机是否存在以及是否与在。
func (channelManager *ChannelManager) ExchangeDeclarePassiveSafe(
	name string, // 交换机名称
	kind string, // 交换机类型
	durable bool, // 是否持久化
	autoDelete bool, // 是否自动删除
	internal bool, // 是否为内部交换机
	noWait bool, // 是否阻塞处理
	args amqp.Table, // 额外的属性
) error {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	return channelManager.ch.ExchangeDeclarePassive(name, kind, durable, autoDelete, internal, noWait, args)
}

// QueueDeclareSafe 用于声明一个队列。如果队列不存在，会创建一个新的队列
// 如果队列不存在，服务器会创建一个新的队列。如果队列已经存在且参数与现有队列的参数不同，会返回一个错误。
func (channelManager *ChannelManager) QueueDeclareSafe(
	name string, //队列名称
	durable bool, //队列是否持久化
	autoDelete bool, //是否自动删除
	exclusive bool, //是否具有排他性
	noWait bool, //是否阻塞处理
	args amqp.Table, //额外的属性
) (string, error) {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	quene, err := channelManager.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return "", err
	}

	return quene.Name, nil
}

// QueueDeclarePassiveSafe 用于声明一个队列。如果队列不存在，会创建一个新的队列
// 不会创建新的队列。仅用于检查队列是否存在以及是否与给定参数匹配。如果队列不存在或者参数不匹配，会返回一个错误
func (channelManager *ChannelManager) QueueDeclarePassiveSafe(
	name string, //队列名称
	durable bool, //队列是否持久化
	autoDelete bool, //是否自动删除
	exclusive bool, //是否具有排他性
	noWait bool, //是否阻塞处理
	args amqp.Table, //额外的属性
) (string, error) {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	quene, err := channelManager.ch.QueueDeclarePassive(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return "", err
	}

	return quene.Name, nil
}

func (channelManager *ChannelManager) PublishWithContextSafe(
	ctx context.Context,
	exchange string, //交换机名称
	key string, //转发的路由key
	mandatory bool, //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
	immediate bool, //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
	msg amqp.Publishing,
) error {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	return channelManager.ch.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
}
func (channelManager *ChannelManager) ConsumeWithContextSafe(
	ctx context.Context,
	name string, // 队列名称，从该队列中消费消息
	consumer string, //消费者标签，用于区分不同的消费者
	autoAck bool, //自动确认消息，如果为 true，则消息被消费后自动确认
	exclusive bool, //是否独占队列，如果为 true，则该队列只能被此消费者消费
	noLocal bool, //如果为 true，则服务器不会将发布到同一连接的消息传递给这个消费者
	noWait bool, //是否阻塞处理，如果为 true，声明队列时不会等待服务器的响应
	args amqp.Table, //额外的属性，用于扩展功能
) (<-chan amqp.Delivery, error) {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()
	return channelManager.ch.ConsumeWithContext(ctx, name, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (channelManager *ChannelManager) ConsumeQueneBindSafe(
	name string, // 队列名称， 要绑定的队列名称。
	key string, // 指定消息的路由键。交换机会根据这个键来决定将消息发送到哪个队列。
	exchange string, //指定将要绑定的交换机名称。消息将从这个交换机路由到绑定的队列。
	noWait bool, //如果设置为 true，表示不等待服务器的响应。绑定操作立即返回，不会确认绑定是否成功。
	args amqp.Table, //传递一些额外的绑定参数，可以用于配置更复杂的绑定规则或扩展功能。例如，可以设置绑定的优先级等。
) error {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	return channelManager.ch.QueueBind(name, key, exchange, noWait, args)
}
