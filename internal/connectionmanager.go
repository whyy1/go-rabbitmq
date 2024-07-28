package internal

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type ConnectionManager struct {
	source      string
	conn        *amqp.Connection
	connLocker  *sync.RWMutex
	stopSignal  chan struct{} //通知全局Coon停止重连机制
	options     ConnectionOptions
	channelPool sync.Pool
}

func NewCoon(url string, opts ...func(*ConnectionOptions)) (connectionManager *ConnectionManager, err error) {
	connectionManager = &ConnectionManager{
		source:     url,
		connLocker: &sync.RWMutex{},
		stopSignal: make(chan struct{}, 1),
		options:    setDefaultConnectionOptions(), //设置默认的链接参数

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

	go connectionManager.startNotifyClose()

	return connectionManager, nil
}

// 启动错误监听
func (connectionManager *ConnectionManager) startNotifyClose() {
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

// 重连Coon循环
func (connectionManager *ConnectionManager) reconnectCoonLoop() {
	for {
		connectionManager.connLocker.Lock()

		select {
		case <-connectionManager.stopSignal:
			connectionManager.options.Logger.Infof("收到全局Coon停止，停止重连循环队列")

			connectionManager.connLocker.Unlock()
			return
		default:
			if err := connectionManager.reconnect(); err != nil {
				connectionManager.options.Logger.Errorf("Connection reconnect err=%v", err)
				connectionManager.connLocker.Unlock()
				continue
			}
			//重新连接成功则启动监听器
			go connectionManager.startNotifyClose()

			connectionManager.connLocker.Unlock()
			return
		}
	}
}

func (connectionManager *ConnectionManager) reconnect() (err error) {
	//等待指定时间后重新连接
	connectionManager.options.Logger.Infof("%v Start reconnecting", connectionManager.options.ReconnectInterval)
	time.Sleep(connectionManager.options.ReconnectInterval)
	connectionManager.options.Logger.Infof("Start reconnecting")
	newCoon, err := amqp.Dial(connectionManager.source)
	if err != nil {
		connectionManager.options.Logger.Errorf("Reconnecting Dial err=%v", err)
		return err
	}
	connectionManager.options.Logger.Infof("Reconnecting successfully")

	if !connectionManager.conn.IsClosed() {
		//如果原有连接未关闭，关闭原有连接再重新赋值，避免并发问题或资源未释放
		if err := connectionManager.conn.Close(); err != nil {
			connectionManager.options.Logger.Errorf("Old Connection Close err=%v", err)
		}
	}

	connectionManager.conn = newCoon
	connectionManager.clearChannelPool() //清空Channel协程池
	return
}
func (connectionManager *ConnectionManager) clearChannelPool() {
	connectionManager.channelPool = sync.Pool{
		New: func() interface{} {
			connectionManager.connLocker.Lock()
			defer connectionManager.connLocker.Unlock()
			channel, err := connectionManager.conn.Channel()
			if err != nil {
				connectionManager.options.Logger.Errorf("coon新建channel失败,错误为%v", err)
				return nil
			}
			return channel
		},
	}
}

// Pool方式
func newChannelPool(connectionManager *ConnectionManager) func() interface{} {
	return func() interface{} {
		channel, err := connectionManager.conn.Channel()
		if err != nil {
			connectionManager.options.Logger.Errorf("failed to create new ch: %v", err)
			return nil
		}
		return channel
	}
}

//// 正常创建方法
//func (connectionManager *ConnectionManager) newChannel() (*amqp.Channel, error) {
//	return connectionManager.conn.Channel()
//}

func (connectionManager *ConnectionManager) CoonClose() {
	//加锁保证即使正常关闭的时候，也要等上一次重连结束之后再关闭
	connectionManager.connLocker.Lock()
	defer connectionManager.connLocker.Unlock()
	connectionManager.options.Logger.Errorf("coon主动关闭")
	connectionManager.stopSignal <- struct{}{}
	close(connectionManager.stopSignal)
	if err := connectionManager.conn.Close(); err != nil {
		connectionManager.options.Logger.Errorf("coon关闭失败 %v", err)
	}
}
