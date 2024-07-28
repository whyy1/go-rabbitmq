package main

import (
	"context"
	"fmt"
	"github.com/whyy1/go-rabbitmq-pool/internal"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestNewPubilsh(t *testing.T) {
	coon, err := internal.NewCoon(os.Getenv("rabbitmqsource"))
	if err != nil {
		fmt.Println(err)
		return
	}

	pubilsh, err := NewPubilsh(coon)
	if err != nil {
		fmt.Println(err)
		return
	}
	pubilsh2, err := NewPubilsh(coon)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("新建channel成功")
	time.Sleep(10 * time.Second)

	if err := pubilsh.PublishWithContext(context.TODO(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("pubilsh1 发送失败", err)
	}

	if err := pubilsh2.PublishWithContext(context.TODO(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("pubilsh2 发送失败", err)
	}
	fmt.Println("第一次发送结束")
	time.Sleep(10 * time.Second)
	pubilsh.Close()
	fmt.Println("pubilsh已关闭")
	time.Sleep(10 * time.Second)

	if err := pubilsh.PublishWithContext(context.TODO(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("pubilsh1 发送失败", err)
	}

	if err := pubilsh2.PublishWithContext(context.TODO(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
		fmt.Println("pubilsh2 发送失败", err)
	}
	for {
		time.Sleep(3 * time.Second)
		if err := pubilsh.PublishWithContext(context.TODO(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
			fmt.Println("pubilsh1 发送失败", err)
		}
		if err := pubilsh2.PublishWithContext(context.TODO(), []byte("test"), []string{"quene1", "quene2", "quene3"}); err != nil {
			fmt.Println("pubilsh2 发送失败", err)
		}
	}

	//todo 检测Close函数后对象是否回收
	//运行垃圾回收
	runtime.GC()
	fmt.Println("运行结束")

}
