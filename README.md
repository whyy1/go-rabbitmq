# go-rabbitmq-pool

`go-rabbitmq-pool` 是一个用于封装 RabbitMQ 操作的 Go 语言库，提供了连接和通道的重连与复用功能，旨在提高 RabbitMQ 的使用效率和稳定性。

## 已实现功能

1. **Connection 的重连复用**
    - 实现了连接的自动重连机制，确保在连接丢失后可以重新建立连接并继续操作。优化主动关闭时和正处于重连情况的处理逻辑。

2. **Logger 对象的复用**
    - 提供了 Logger 对象的复用功能，以便在不同的操作中保持一致的日志记录。

3. **Channel 的重连复用**
    - 实现了通道的自动重连机制，支持在通道断开后重新建立通道并继续使用。

4. **基础 Publish 的 Demo**
    - 提供了一个简单的示例，展示如何使用复用的通道来发布消息。支持直连和路由两种方式。
   

## TODO

1. **Quene的重试创建持久化**
    - 实现队列的持久化，并在创建队列失败时进行重试，以确保队列的可靠性。

2. **完善 Publish 的发送 Demo**
    - 继续完善发送消息的示例代码，实现直连和路由两种消息发送方式。

3. **实现 Consume 的消费 Demo**
   - 提供消费消息的示例代码，并实现相关参数配置，包括是否需要在释放时重置通道。
   
4. **Channel是否需要来维护sync.Pool**
   - 如果继续按这个方法实现，要将channel重连状态单独为，在ConnectionManager上层统一维护所有channel状态

## 使用说明

### 安装

使用 `go get` 安装 `go-rabbitmq-pool` 库：

```sh

go get github.com/whyy1/go-rabbitmq-pool
