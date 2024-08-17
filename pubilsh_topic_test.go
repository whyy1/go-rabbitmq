package main

import (
	"context"
	"fmt"
	"github.com/whyy1/go-rabbitmq/internal"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestPubilshTopic(t *testing.T) {

	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		fmt.Println(err)
		return
	}

	publisher, err := NewPublisher(coon,
		WithPublisherOptionsExchangeName("topic"),
		WithPublisherOptionsExchangeKind("topic"),
		WithPublisherOptionsExchangeDurable(true),
		WithPublisherOptionsExchangeDeclare(true),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	publisher2, err := NewPublisher(coon,
		WithPublisherOptionsExchangeName("topic"),
		WithPublisherOptionsExchangeKind("topic"),
		WithPublisherOptionsExchangeDurable(true),
		WithPublisherOptionsExchangeDeclare(true),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	sign := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		<-sign
		done <- struct{}{}
	}()
	go func() {
		func() {
			for {
				select {
				case <-ticker.C:
					if err := publisher.PublishWithContext(context.Background(),
						[]byte("topic.test.err"), []string{"topic.test.err"},
						WithPublishOptionsExchange("topic"),
					); err != nil {
						fmt.Println("pubilsh1 发送失败", err)
					}
					log.Println("pubilsh1发送成功")
				}
			}
		}()
	}()

	go func() {
		func() {
			for {
				select {
				case <-ticker.C:
					if err := publisher2.PublishWithContext(context.Background(),
						[]byte("topic.test.log"), []string{"topic.test.log"},
						WithPublishOptionsExchange("topic")); err != nil {
						fmt.Println("publisher2 发送失败", err)
					}
					log.Println("pubilsh2发送成功")
				}
			}
		}()
	}()

	<-done
}
