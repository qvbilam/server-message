package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"message/business"
	"message/global"
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
		zap.S().Fatalf("%s dial error: %s", "队列服务", err)
	}

	global.MessageQueueClient = conn

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
