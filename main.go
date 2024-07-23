package main

import (
	"fmt"
	"go_rabbitmq_pool/internal"
	"log"
	"os"
	"time"
)

var (
	// 连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng
	// 这个信息是固定不变的amqp://是固定参数后面两个是用户名密码ip地址端口号Virtual Host
	source = os.Getenv("rabbitmqsource")
)

func main() {
	//os.Getenv()
	coon, err := internal.NewCoon(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	//coon.CoonClose()

	manager, err := internal.NewChannelManager(coon)
	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)
	manager.ChannelClose()

	_, err = internal.NewChannelManager(coon)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = internal.NewChannelManager(coon)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	//forever := make(chan bool)
	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever
}
