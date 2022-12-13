package business

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"message/global"
	"message/model"
)

type ExchangeBusiness struct {
	Id     int64
	Name   string
	Status int64
}

func (b *ExchangeBusiness) Create() (int64, error) {
	if b.Name == "" {
		b.Name = global.ServerConfig.RabbitMQServerConfig.Exchange
	}
	id := b.GetIdByName()
	if id != 0 { // 交换机存在
		return id, nil
	}
	m := model.QueueExchange{Name: b.Name}
	tx := global.DB.Begin()
	global.DB.Save(&m)
	if err := CreateExchange(b.Name); err != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建交换机失败")
	}
	tx.Commit()
	return m.ID, nil
}

func (b *ExchangeBusiness) GetIdByName() int64 {
	entity := model.QueueExchange{}
	res := global.DB.Where(model.QueueExchange{Name: b.Name}).Select("id").First(&entity)
	if res.RowsAffected == 0 {
		return 0
	}
	return entity.ID
}

func CreateExchange(exchangeName string) error {
	// 建立 amqp 通道
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "建立通道失败", err)
		return err
	}

	// 创建交换机(不存在创建)
	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列交换机", err)
		return err
	}
	return nil
}

// PushExchange 发送交换机消息
func PushExchange(exchangeName string, body []byte) error {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		return err
	}
	if err := ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			Body: body,
		},
	); err != nil {
		return err
	}
	return nil
}

func PushDefaultExchange(body []byte) error {
	exchange := global.ServerConfig.RabbitMQServerConfig.Exchange
	return PushExchange(exchange, body)
}
