package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"message/business"
	"message/global"
	"time"
)

func InitQueue() {
	//user := "admin"
	//password := "admin"
	//host := "127.0.0.1"
	//port := 5672

	host := global.ServerConfig.RabbitMQServerConfig.Host
	port := global.ServerConfig.RabbitMQServerConfig.Port
	user := global.ServerConfig.RabbitMQServerConfig.User
	password := global.ServerConfig.RabbitMQServerConfig.Password

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		//zap.S().Fatalf("%s dial error: %s", "队列服务", err)
		zap.S().Errorf("%s dial error: %s", "队列服务", err)
	}

	global.MessageQueueClient = conn
	if global.MessageQueueClient.IsClosed() {
		return
	}

	fmt.Printf("create queue exchange: %s\n", global.ServerConfig.RabbitMQServerConfig.Exchange)

	// 创建消息队列
	_ = business.CreateExchange(global.ServerConfig.RabbitMQServerConfig.Exchange)
	// 创建用户聊天队列
	_ = business.CreateExchange(global.ServerConfig.RabbitMQServerConfig.ExchangeChatPrivate)
	// 创建群组聊天队列
	_ = business.CreateExchange(global.ServerConfig.RabbitMQServerConfig.ExchangeChatGroup)
	// 创建房间聊天队列
	_ = business.CreateExchange(global.ServerConfig.RabbitMQServerConfig.ExchangeChatRoom)
}

func InitQueueHealth() {

	// 监听状态
	closeChan := make(chan *amqp.Error, 1)
	notifyClose := global.MessageQueueClient.NotifyClose(closeChan)

	// 健康检查
	timer := time.NewTimer(5 * time.Second)

	for {
		select {
		case e := <-notifyClose:
			if e != nil {
				fmt.Printf("chan通道错误,e:%+v\n", e)
				InitQueue()
			}
			//close(closeChan)
			//InitQueue()
		case <-timer.C:
			//timer.Reset(time.Second * time.Duration(rand.Intn(5)))
			timer.Reset(time.Second * 5)
			if global.MessageQueueClient.IsClosed() == true {
				fmt.Printf("定期检查错误，尝试重启\n")
				InitQueue()
			}
			//fmt.Printf("定时检测rabbitMq: %d\n", time.Now().Second())
		}
	}

}
