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

// ExchangeDeclarePassiveSafe 用于检查交换机是否存在。如果交换机不存在，会返回一个错误
// 不会创建新的交换机。仅用于检查交换机是否存在以及是否与给定参数匹配。如果交换机不存在或者参数不匹配，会返回一个错误
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
) error {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	_, err := channelManager.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return err
	}

	return err
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
) error {
	channelManager.chLocker.RLock()
	defer channelManager.chLocker.RUnlock()

	_, err := channelManager.ch.QueueDeclarePassive(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return err
	}

	return err
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
