package internal

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type ChannelManager struct {
	connectionManager  *ConnectionManager
	channel            *amqp.Channel
	options            ConnectionOptions
	channelNotifyClose chan *amqp.Error
}

func NewChannelManager(connectionManager *ConnectionManager) (channelManager *ChannelManager, err error) {
	channelManager = &ChannelManager{
		connectionManager: connectionManager,
		options:           connectionManager.options,
	}
	channel := getNewChannelPool(connectionManager)
	if channel == nil {
		return nil, errors.New("getNewChannelPool is null")
	}
	channelManager.channel = channel
	go channelManager.StartNotifyClose()

	return channelManager, nil
}

func (channelManager *ChannelManager) StartNotifyClose() {
	channelNotifyClose := channelManager.channel.NotifyClose(make(chan *amqp.Error, 1))
	err := <-channelNotifyClose
	if err != nil {
		//非正常关闭
		channelManager.options.Logger.Errorf("channelNotifyClose err=%v", err)
		channelManager.reconnectCoonLoop()
		return
	}
	channelManager.options.Logger.Infof("Channel closed normally.")
}

func (channelManager *ChannelManager) reconnectCoonLoop() {
	for {

		if err := channelManager.reconnect(); err != nil {
			channelManager.options.Logger.Errorf("Connection reconnect err=%v", err)
			continue
		}
		//重新连接成功则启动监听器
		channelManager.channelNotifyClose = channelManager.channel.NotifyClose(make(chan *amqp.Error, 1))
		go channelManager.StartNotifyClose()
		return
	}
}

func (channelManager *ChannelManager) reconnect() (err error) {
	//等待指定时间后重新连接
	channelManager.options.Logger.Infof("%v Start reconnecting", channelManager.options.ReconnectInterval)
	time.Sleep(channelManager.options.ReconnectInterval)
	channelManager.options.Logger.Infof("Start reconnecting")
	channel := getNewChannelPool(channelManager.connectionManager)
	if channel == nil {
		channelManager.options.Logger.Errorf("getNewChannelPool is null")
		return err
	}
	channelManager.options.Logger.Infof("Reconnecting successfully")
	//关闭现有连接再重新赋值，避免并发问题或资源未释放
	if err := channelManager.channel.Close(); err != nil {
		channelManager.options.Logger.Errorf("Old Channel Close err=%v", err)
	}
	channelManager.channel = channel
	return
}
func (channelManager *ChannelManager) ChannelClose() {
	//归还Channel
	putNewChannelPool(channelManager.connectionManager, channelManager.channel)
}

func getNewChannelPool(connectionManager *ConnectionManager) *amqp.Channel {
	//connectionManager.connLocker.Lock()
	//defer connectionManager.connLocker.Unlock()
	channel := connectionManager.channelPool.Get()
	if channel == nil {
		return nil
	}
	return channel.(*amqp.Channel)
}
func putNewChannelPool(connectionManager *ConnectionManager, channel *amqp.Channel) {
	//connectionManager.connLocker.Lock()
	//defer connectionManager.connLocker.Unlock()
	connectionManager.channelPool.Put(channel)
}
