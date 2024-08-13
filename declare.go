package main

import (
	"github.com/whyy1/go-rabbitmq/internal"
)

func declareExchange(channelManager *internal.ChannelManager, options ExchangeOptions) error {
	if !options.Declare {
		return nil
	}

	if options.Passive {

		return channelManager.ExchangeDeclarePassiveSafe(
			options.Name,
			options.Kind,
			options.Durable,
			options.AutoDelete,
			options.Internal,
			options.NoWait,
			options.Args,
		)
	}

	return channelManager.ExchangeDeclareSafe(
		options.Name,
		options.Kind,
		options.Durable,
		options.AutoDelete,
		options.Internal,
		options.NoWait,
		options.Args,
	)
}
func declareQuene(channelManager *internal.ChannelManager, options *QueueOptions) error {
	if !options.Declare {
		return nil
	}

	if options.Passive {

		queneName, err := channelManager.QueueDeclarePassiveSafe(
			options.Name,
			options.Durable,
			options.AutoDelete,
			options.Exclusive,
			options.NoWait,
			options.Args,
		)
		if err != nil {
			return err
		} else {
			options.Name = queneName
			return nil
		}
	}

	queneName, err := channelManager.QueueDeclareSafe(
		options.Name,
		options.Durable,
		options.AutoDelete,
		options.Exclusive,
		options.NoWait,
		options.Args,
	)
	if err != nil {
		return err
	} else {
		options.Name = queneName
		return nil
	}
}

func queneBind(channelManager *internal.ChannelManager, queueBindOptions QueueBindOptions) error {
	if !queueBindOptions.Bind {
		return nil
	}

	return channelManager.ConsumeQueneBindSafe(
		queueBindOptions.Name,
		queueBindOptions.Key,
		queueBindOptions.Exchange,
		queueBindOptions.NoWait,
		queueBindOptions.Args,
	)
}
