package main

import amqp "github.com/rabbitmq/amqp091-go"

type ConsumOptions struct {
	Name      string
	AutoAck   bool
	Exclusive bool
	NoWait    bool
	NoLocal   bool
	Args      amqp.Table
}
