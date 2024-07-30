package main

import (
	"github.com/whyy1/go-rabbitmq-pool/internal"
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
func declareQuene(channelManager *internal.ChannelManager, options QueueOptions) error {
	if !options.Declare {
		return nil
	}

	if options.Passive {

		return channelManager.QueueDeclarePassiveSafe(
			options.Name,
			options.Durable,
			options.AutoDelete,
			options.Exclusive,
			options.NoWait,
			options.Args,
		)
	}

	return channelManager.QueueDeclareSafe(
		options.Name,
		options.Durable,
		options.AutoDelete,
		options.Exclusive,
		options.NoWait,
		options.Args,
	)
}
