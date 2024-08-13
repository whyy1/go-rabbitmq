package main

import (
	"context"
	"github.com/whyy1/go-rabbitmq/internal"
	"log"
	"os"
	"testing"
	"time"
)

func TestConsumerSimple(t *testing.T) {
	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		log.Println(err)
		return
	}
	consumer, err := NewConsumer(coon,
		WithConsumerQueueName("quene1"),
		WithConsumerQueueDeclare(true))
	if err != nil {
		log.Println(err)
		return
	}

	for {
		channel, err := consumer.GetConsumeChannel(context.Background(),
			WithConsumeAutoAck(true),
		)
		if err != nil {
			log.Println(err)
			return
		}

		for delivery := range channel {
			log.Println("接收到消息", string(delivery.Body))

		}
		time.Sleep(6 * time.Second)
		log.Println("channel close")
	}

}
