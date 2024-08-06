package main

import amqp "github.com/rabbitmq/amqp091-go"

type ConsumeOptions struct {
	Name      string
	AutoAck   bool
	Exclusive bool
	NoWait    bool
	NoLocal   bool
	Args      amqp.Table
}

func WithConsumeName(name string) func(options *ConsumeOptions) {
	return func(options *ConsumeOptions) {
		options.Name = name
	}
}

func WithConsumeAutoAck(autoAck bool) func(options *ConsumeOptions) {
	return func(options *ConsumeOptions) {
		options.AutoAck = autoAck
	}
}
func WithConsumeExclusive(exclusive bool) func(options *ConsumeOptions) {
	return func(options *ConsumeOptions) {
		options.Exclusive = exclusive
	}
}

func WithConsumeNoWait(noWait bool) func(options *ConsumeOptions) {
	return func(options *ConsumeOptions) {
		options.NoWait = noWait
	}
}

func WithConsumeNoLocal(noLocal bool) func(options *ConsumeOptions) {
	return func(options *ConsumeOptions) {
		options.NoLocal = noLocal
	}
}
func WithConsumeArgs(args amqp.Table) func(options *ConsumeOptions) {
	return func(options *ConsumeOptions) {
		options.Args = args
	}
}
