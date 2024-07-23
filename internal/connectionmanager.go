package internal

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type ConnectionManager struct {
	source          string
	conn            *amqp.Connection
	connLocker      sync.Locker
	options         ConnectionOptions
	channelPool     sync.Pool
	connNotifyClose chan *amqp.Error
}

func NewCoon(url string, opts ...func(*ConnectionOptions)) (connectionManager *ConnectionManager, err error) {
	connectionManager = &ConnectionManager{
		source:  url,
		options: SetDefaultConnectionOptions(), //设置默认的链接参数
	}
	//设置参数
	for _, optFunc := range opts {
		optFunc(&connectionManager.options)
	}
	coon, err := amqp.Dial(connectionManager.source)
	if err != nil {
		connectionManager.options.Logger.Errorf("Dial err%v", err)
		return connectionManager, err
	}
	connectionManager.conn = coon
	// 初始化 channelPool 的 New 函数
	connectionManager.channelPool.New = newChannelPool(connectionManager)

	go connectionManager.StartNotifyClose()

	return connectionManager, nil
}
func (connectionManager *ConnectionManager) StartNotifyClose() {
	connNotifyClose := connectionManager.conn.NotifyClose(make(chan *amqp.Error, 1))
	err := <-connNotifyClose
	if err != nil {
		//非正常关闭
		connectionManager.options.Logger.Errorf("connNotifyClose err=%v", err)
		connectionManager.reconnectCoonLoop()
		return
	}
	connectionManager.options.Logger.Infof("Coon closed normally.")
}

func (connectionManager *ConnectionManager) reconnectCoonLoop() {
	for {

		if err := connectionManager.reconnect(); err != nil {
			connectionManager.options.Logger.Errorf("Connection reconnect err=%v", err)
			continue
		}
		//重新连接成功则启动监听器
		connectionManager.connNotifyClose = connectionManager.conn.NotifyClose(make(chan *amqp.Error, 1))
		go connectionManager.StartNotifyClose()
		return
	}
}
func (connectionManager *ConnectionManager) reconnect() (err error) {
	//等待指定时间后重新连接
	connectionManager.options.Logger.Infof("%v Start reconnecting", connectionManager.options.ReconnectInterval)
	time.Sleep(connectionManager.options.ReconnectInterval)
	connectionManager.options.Logger.Infof("Start reconnecting")
	coon, err := amqp.Dial(connectionManager.source)
	if err != nil {
		connectionManager.options.Logger.Errorf("Reconnecting Dial err=%v", err)
		return err
	}
	connectionManager.options.Logger.Infof("Reconnecting successfully")
	//关闭现有连接再重新赋值，避免并发问题或资源未释放
	if err := connectionManager.conn.Close(); err != nil {
		connectionManager.options.Logger.Errorf("Old Connection Close err=%v", err)
	}
	connectionManager.conn = coon
	// 初始化 channelPool 的 New 函数
	connectionManager.channelPool.New = newChannelPool(connectionManager)
	return
}
func newChannelPool(connectionManager *ConnectionManager) func() interface{} {
	return func() interface{} {
		channel, err := connectionManager.conn.Channel()
		if err != nil {
			connectionManager.options.Logger.Errorf("failed to create new channel: %v", err)
			return nil
		}
		return channel
	}
}

func (connectionManager *ConnectionManager) CoonClose() {
	connectionManager.conn.Close()
	//todo 需要关闭该connection下所有的Channel
}
