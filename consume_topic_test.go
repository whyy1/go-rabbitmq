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

func TestConsumerTopic(t *testing.T) {
	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			func() {
				key := "#"
				consumer, err := NewConsumer(coon,
					WithConsumerQueueAutoDelete(true),
					WithConsumerQueueDeclare(true),
					WithConsumerExchangeName("topic"),
					WithConsumerBindKey(key),
					WithConsumerBind(true),
				)
				defer consumer.Close()
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
					log.Println(key+"接收到消息", string(delivery.Body))
				}
			}()
		}
	}()
	go func() {
		for {
			func() {
				key := "*.err"
				consumer, err := NewConsumer(coon,
					WithConsumerQueueAutoDelete(true),
					WithConsumerQueueDeclare(true),
					WithConsumerExchangeName("topic"),
					WithConsumerBindKey(key),
					WithConsumerBind(true),
				)
				defer consumer.Close()
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
					log.Println(key+"接收到消息", string(delivery.Body))
				}

				log.Println("channel结束")
			}()
		}
	}()

	go func() {
		for {
			func() {
				key := "topic.*.log"
				consumer, err := NewConsumer(coon,
					WithConsumerQueueAutoDelete(true),
					WithConsumerQueueDeclare(true),
					WithConsumerExchangeName("topic"),
					WithConsumerBindKey(key),
					WithConsumerBind(true),
				)
				defer consumer.Close()
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
					log.Println(key+"接收到消息", string(delivery.Body))
				}

				log.Println("channel结束")
			}()
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	<-sign
}
