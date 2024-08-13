package main

import (
	"context"
	"github.com/whyy1/go-rabbitmq/internal"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestConsumerPublish(t *testing.T) {
	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		for {
			consumer, err := NewConsumer(coon,
				WithConsumerQueueDeclare(true),
				WithConsumerQueueAutoDelete(true),
				WithConsumerExchangeName("amq.fanout"),
				WithConsumerQueueBind(true),
			)
			if err != nil {
				log.Println("新建NewConsumer错误", err)
				return
			}

			channel, err := consumer.GetConsumeChannel(context.Background(),
				WithConsumeAutoAck(true),
			)
			if err != nil {
				log.Println("新建GetConsumeChannel错误", err)
				return
			}

			for delivery := range channel {
				log.Println("Consumer1接收到消息", string(delivery.Body))

			}
		}
	}()
	go func() {
		for {
			consumer, err := NewConsumer(coon,
				WithConsumerQueueDeclare(true),
				WithConsumerQueueAutoDelete(true),
				WithConsumerExchangeName("amq.fanout"),
				WithConsumerQueueBind(true),
			)
			if err != nil {
				log.Println("新建NewConsumer错误", err)
				return
			}

			channel, err := consumer.GetConsumeChannel(context.Background(),
				WithConsumeAutoAck(true),
			)
			if err != nil {
				log.Println("新建GetConsumeChannel错误", err)
				return
			}

			for delivery := range channel {
				log.Println("Consumer2接收到消息", string(delivery.Body))
			}
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	<-sign
}
