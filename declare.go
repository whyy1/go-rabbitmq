package main

import "github.com/whyy1/go-rabbitmq-pool/internal"

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
