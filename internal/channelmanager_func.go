package internal

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

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
