package main

//
//// ExchangeOptions are used to configure an exchange.
//// If the Passive flag is set the client will only check if the exchange exists on the server
//// and that the settings match, no creation attempt will be made.
//type ExchangeOptions struct {
//	Name       string
//	Kind       string // possible values: empty string for default exchange or direct, topic, fanout
//	Durable    bool
//	AutoDelete bool
//	Internal   bool
//	NoWait     bool
//	Passive    bool // if false, a missing exchange will be created on the server
//	//Args       Table
//	Declare  bool
//	Bindings []string
//}
//type PublishOptions struct {
//	ReconnectInterval time.Duration
//	Logger            internal.Logger
//	ExchangeOptions   ExchangeOptions
//}
//
//func WithDefaultPublishOptionsOptions(connectionOptions *internal.ConnectionOptions) (options PublishOptions) {
//	return PublishOptions{
//		ReconnectInterval: connectionOptions.ReconnectInterval,
//		Logger:            connectionOptions.Logger,
//		ExchangeOptions: ExchangeOptions{
//			Name:       "",
//			Kind:       amqp.ExchangeDirect,
//			Durable:    false,
//			AutoDelete: false,
//			Internal:   false,
//			NoWait:     false,
//			Passive:    false,
//			//Args:       Table{},
//			Declare: false,
//		},
//	}
//}
//
//func WithReconnectInterval(interval time.Duration) func(options *PublishOptions) {
//	return func(options *PublishOptions) {
//		options.ReconnectInterval = interval
//	}
//}
//func WithLogger(logger internal.Logger) func(options *PublishOptions) {
//	return func(options *PublishOptions) {
//		options.Logger = logger
//	}
//}
