package internal

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

const (
	ChannelNormalShutdown     ChannelManagerError = "ChannelNormalShutdown"
	ChannelUnexpectedShutdown ChannelManagerError = "ChannelUnexpectedShutdown"
	ManagerClose              ChannelManagerError = "ManagerClose"
)

type ChannelManagerError string
type ChannelManager struct {
	connectionManager *ConnectionManager
	ch                *amqp.Channel
	chLocker          *sync.RWMutex
	stopSignal        chan struct{} //通知全局Channel停止重连机制
	options           ConnectionOptions
}

func NewChannelManager(connectionManager *ConnectionManager) (channelManager *ChannelManager, err error) {
	channelManager = &ChannelManager{
		connectionManager: connectionManager,
		chLocker:          &sync.RWMutex{},
		stopSignal:        make(chan struct{}, 1),
		options:           connectionManager.options,
	}

	newChannel := channelManager.getChannel()
	if newChannel == nil {
		return nil, errors.New("getChannel is null")
	}
	channelManager.ch = newChannel
	go channelManager.startNotifyClose()

	return channelManager, nil
}

func (channelManager *ChannelManager) PublishWithContext(exchange string, key string, mandatory bool, immediate bool, data []byte) (err error) {
	channelManager.ch.QueueDeclare(key, true, false, false, false, nil)
	return channelManager.ch.Publish(exchange, key, mandatory, immediate, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         data,
	})
}

// StartNotifyClose 监听Channel是否成功关闭，非正常关闭则重新连接
func (channelManager *ChannelManager) startNotifyClose() {
	channelNotifyClose := channelManager.ch.NotifyClose(make(chan *amqp.Error, 1))

	err := <-channelNotifyClose
	if err != nil {
		//非正常关闭,通知Manager尝试重连
		channelManager.options.Logger.Errorf("Channel接收到意外关闭%v", err)
		channelManager.reconnectCoonLoop()
		return
	}

	channelManager.options.Logger.Infof("Channel正常关闭")

}

// 监听Manager状态，判断是否需要尝试重新连接Channel
//func (channelManager *ChannelManager) startNotifyManagerClose() {
//	for {
//		select {
//		case err := <-channelManager.channelManagerNotifyClose:
//			switch err {
//			case ChannelNormalShutdown: //channel被正常关闭
//				channelManager.options.Logger.Infof("channel被正常关闭")
//				return
//			case ChannelUnexpectedShutdown:
//				channelManager.reconnectCoonLoop()
//			case ManagerClose: //正常关闭channelManager
//				channelManager.options.Logger.Infof("正常关闭channelManager被正常关闭")
//				return
//			}
//		}
//	}
//	//for {
//	//	select {
//	//	case err := <-channelNotifyClose:
//	//		{
//	//			if err != nil {
//	//				channelManager.options.Logger.Errorf("Channel接收到非正常关闭%v", err)
//	//				//非正常关闭
//	//				channelManager.options.Logger.Errorf("channelNotifyClose err=%v", err)
//	//				channelManager.reconnectCoonLoop()
//	//				channelManager.channelManagerNotifyClose <- ChannelUnexpectedShutdown
//	//				//return
//	//			}
//	//			channelManager.channelManagerNotifyClose <- ChannelNormalShutdown
//	//			channelManager.options.Logger.Infof("Channel closed normally.")
//	//		}
//	//	case err := <-channelManager.channelManagerNotifyClose:
//	//		{
//	//			switch err {
//	//			case ChannelNormalShutdown: //channel被正常关闭
//	//				channelManager.options.Logger.Infof("channel被正常关闭")
//	//				//return
//	//			case ChannelUnexpectedShutdown:
//	//				for {
//	//					newChannel := getNewChannelPool(channelManager.connectionManager)
//	//					if newChannel == nil {
//	//						channelManager.options.Logger.Infof("channelManager接收重连自身channel失败，%v后重试", channelManager.options.ReconnectInterval)
//	//						time.Sleep(channelManager.options.ReconnectInterval)
//	//					}
//	//					channelManager.options.Logger.Infof("channelManager接收重连自身channel成功")
//	//					channelManager.ch = newChannel
//	//					break
//	//				}
//	//			case ManagerClose: //正常关闭channelManager
//	//				channelManager.options.Logger.Infof("正常关闭channelManager被正常关闭")
//	//				//return
//	//			}
//	//		}
//	//	}
//	//}
//}

// 不断循环尝试重新连接Channel
func (channelManager *ChannelManager) reconnectCoonLoop() {
	for {
		channelManager.chLocker.Lock()
		select {
		case <-channelManager.stopSignal:
			channelManager.options.Logger.Infof("收到Channel停止，停止重连循环队列")

			channelManager.chLocker.Unlock()
			return
		default:
			if err := channelManager.reconnect(); err != nil {
				channelManager.options.Logger.Errorf("重新连接Channel失败,错误为%v", err)

				channelManager.chLocker.Unlock()
				continue
			}
			//重新连接成功则启动监听器
			go channelManager.startNotifyClose()

			channelManager.chLocker.Unlock()
			return
		}
	}
}

// 重新连接Channel
func (channelManager *ChannelManager) reconnect() (err error) {
	//等待指定时间后重新连接
	channelManager.options.Logger.Infof("%v后开始重新连接Channel", channelManager.options.ReconnectInterval)
	time.Sleep(channelManager.options.ReconnectInterval)
	channelManager.options.Logger.Infof("开始重新连接Channel")
	newChannel := channelManager.getChannel()
	if newChannel == nil {
		return errors.New("获取Channel为空")
	}
	channelManager.options.Logger.Infof("重新连接Channel成功")
	//关闭原有连接再重新赋值，避免并发问题或资源未释放
	if err := channelManager.ch.Close(); err != nil {
		channelManager.options.Logger.Errorf("原有Channel关闭错误%v", err)
	}
	channelManager.ch = newChannel
	return
}
func (channelManager *ChannelManager) getChannel() *amqp.Channel {
	channel := channelManager.connectionManager.channelPool.Get()
	if channel == nil {
		return nil
	}

	return channel.(*amqp.Channel)
}

func (channelManager *ChannelManager) putChannel() {
	channelManager.chLocker.Lock()
	defer channelManager.chLocker.Unlock()
	//归还Channel
	channelManager.connectionManager.channelPool.Put(channelManager.ch)
	//todo 是否需要将对象置空避免没有回收导致内存泄露
	//channelManager.ch = nil
	//channelManager.connectionManager = nil
}

func (channelManager *ChannelManager) CloseChannel() {
	channelManager.chLocker.Lock()
	defer channelManager.chLocker.Unlock()
	channelManager.stopSignal <- struct{}{}

	close(channelManager.stopSignal)
	//todo 是否需要将对象置空避免没有回收导致内存泄露
	//将引用对象置空
	//channelManager.ch = nil
	//channelManager.connectionManager = nil
	if err := channelManager.ch.Close(); err != nil {
		channelManager.options.Logger.Errorf("channel关闭失败 %v", err)
	}
}
