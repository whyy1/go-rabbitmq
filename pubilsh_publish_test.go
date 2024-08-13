package main

import (
	"context"
	"fmt"
	"github.com/whyy1/go-rabbitmq/internal"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func TestPubilshPubilsh(t *testing.T) {

	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		fmt.Println(err)
		return
	}

	pubilsh, err := NewPubilsh(coon,
		WithPublisherOptionsExchangeName("amq.fanout"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("新建channel成功")
	//time.Sleep(10 * time.Second)

	if err := pubilsh.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"},
		WithPublishOptionsExchange("amq.fanout"),
	); err != nil {
		fmt.Println("pubilsh1 发送失败", err)
	}

	sign := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		<-sign
		done <- struct{}{}
	}()
	for {
		select {
		case <-ticker.C:
			if err := pubilsh.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"},
				WithPublishOptionsExchange("amq.fanout"),
			); err != nil {
				fmt.Println("pubilsh1 发送失败", err)
			}
			log.Println("pubilsh1发送成功")
		case <-done:
			//todo 检测Close函数后对象是否回收
			//运行垃圾回收
			runtime.GC()
			fmt.Println("运行结束")
			return
		}
	}

}
