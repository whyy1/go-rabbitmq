package main

import (
	"context"
	"fmt"
	"github.com/whyy1/go-rabbitmq/internal"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func TestPubilshSimple(t *testing.T) {
	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		fmt.Println(err)
		return
	}

	publisher, err := NewPublisher(coon)
	if err != nil {
		fmt.Println(err)
		return
	}
	publisher2, err := NewPublisher(coon)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("新建channel成功")
	time.Sleep(10 * time.Second)

	if err := publisher.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("pubilsh1 发送失败", err)
	}

	if err := publisher2.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("publisher2 发送失败", err)
	}
	fmt.Println("第一次发送结束")
	time.Sleep(10 * time.Second)
	publisher.Close()
	fmt.Println("pubilsh已关闭")
	time.Sleep(10 * time.Second)

	if err := publisher.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("pubilsh1 发送失败", err)
	}

	if err := publisher2.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("publisher2 发送失败", err)
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
			//if err := publisher.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
			//	fmt.Println("pubilsh1 发送失败", err)
			//}
			if err := publisher2.PublishWithContext(context.Background(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
				fmt.Println("publisher2 发送失败", err)
			}
		case <-done:
			//todo 检测Close函数后对象是否回收
			//运行垃圾回收
			runtime.GC()
			fmt.Println("运行结束")
			return
		}
	}

}
