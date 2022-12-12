package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"message/business"
	"message/global"
)

func InitQueue() {
	user := "admin"
	password := "admin"
	host := "127.0.0.1"
	port := 5672

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列服务", err)
	}

	global.MessageQueueClient = conn

	ExchangeName := global.ServerConfig.RabbitMQServerConfig.Exchange

	fmt.Printf("create queue exchange: %s\n", ExchangeName)

	// 全局变量
	global.ServerConfig.RabbitMQServerConfig.Exchange = ExchangeName

	// 创建队列
	_ = business.CreateExchange(ExchangeName)
}
