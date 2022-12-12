package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"message/business"
	"message/global"
	"strconv"
)

var ExchangeName = "qvbilam-message-exchange"
var QueueNamePrefix = "qvbilam-message-queue-"

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
	suffix := global.ServerConfig.RabbitMQServerConfig.QueueSuffix
	if suffix == "" {
		suffix = strconv.Itoa(int(global.ServerConfig.Port))
	}
	QueueName := QueueNamePrefix + suffix

	fmt.Printf("create queue exchange: %s\n", ExchangeName)
	fmt.Printf("create queue: %s\n", QueueName)

	// 全局变量
	global.ServerConfig.RabbitMQServerConfig.MessageExchangeName = ExchangeName
	global.ServerConfig.RabbitMQServerConfig.MessageQueueName = QueueName

	// 创建队列
	business.CreateExchange(ExchangeName)
	business.CreateQueue(QueueName, ExchangeName)

	// 接受消息
	go business.ConsumeQueue(QueueName)
}