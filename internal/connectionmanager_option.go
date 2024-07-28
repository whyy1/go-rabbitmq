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

func setDefaultConnectionOptions() ConnectionOptions {
	logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any
	return ConnectionOptions{
		ReconnectInterval: 5 * time.Second,
		Logger:            logger.Sugar(),
	}
}

func SetReconnectInterval(interval time.Duration) func(options *ConnectionOptions) {
	return func(options *ConnectionOptions) {
		options.ReconnectInterval = interval
	}
}
func SetLogger(logger Logger) func(options *ConnectionOptions) {
	return func(options *ConnectionOptions) {
		options.Logger = logger
	}
}

func (connectionManager *ConnectionManager) GetConnectionOptions() ConnectionOptions {
	return connectionManager.options
}
