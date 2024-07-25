package main

import (
	"fmt"
	"github.com/whyy1/go-rabbitmq-pool/internal"
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
	coon, err := internal.NewCoon(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	//_, err = internal.NewChannelManager(coon)
	//if err != nil {
	//	return
	//}

	pubilsh, err := NewPubilsh(coon)
	if err != nil {
		fmt.Println(err)
		return
	}

	pubilsh.PublishWithContext("", "test", false, false, []byte("test"))
	time.Sleep(8 * time.Second)
	pubilsh.Close()

	pubilsh1, err := NewPubilsh(coon)
	if err != nil {
		fmt.Println(err)
		return
	}
	pubilsh1.PublishWithContext("", "test", false, false, []byte("test1"))
	//pubilsh1.Close()
	time.Sleep(2 * time.Second)

	pubilsh2, err := NewPubilsh(coon)
	if err != nil {
		fmt.Println(err)
		return
	}
	pubilsh2.PublishWithContext("", "test", false, false, []byte("test2"))
	for {

	}
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	//forever := make(chan bool)
	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever
}
