package internal

import (
	"go.uber.org/zap"
	"time"
)

type ConnectionOptions struct {
	ReconnectInterval time.Duration
	Logger            Logger
	//Config            Config
}

func withDefaultConnectionOptions() ConnectionOptions {
	logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any
	return ConnectionOptions{
		ReconnectInterval: 5 * time.Second,
		Logger:            logger.Sugar(),
	}
}

func WithConnectionReconnectInterval(interval time.Duration) func(options *ConnectionOptions) {
	return func(options *ConnectionOptions) {
		options.ReconnectInterval = interval
	}
}
func WithConnectionLogger(logger Logger) func(options *ConnectionOptions) {
	return func(options *ConnectionOptions) {
		options.Logger = logger
	}
}
