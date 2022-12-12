package business

import (
	"fmt"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"message/enum"
	"message/global"
	"message/model"
)

type QueueBusiness struct {
	Id           int64
	Name         string
	ExchangeId   int64
	ExchangeName string
	Status       *int64
}

// Create 创建(队列创建为空闲状态, web服务创建队列后调用修改状态)
func (b *QueueBusiness) Create() (*model.Queue, error) {
	// 创建交换机(不存在创建
	eb := ExchangeBusiness{}
	eId, err := eb.Create()
	if err != nil {
		return nil, err
	}

	// 获取空余队列
	b.ExchangeId = eId
	idleEntity, err := b.GetExchangeIdleQueue()
	if err == nil && idleEntity != nil {
		return idleEntity, err
	}

	// 创建队列
	if id := b.GetIdByName(); id != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "队列存在")
	}
	entity := model.Queue{
		Name:       b.Name,
		ExchangeID: eId,
	}
	global.DB.Save(&entity)
	return &entity, nil
}

// GetExchangeIdleQueue 获取空闲队列
func (b *QueueBusiness) GetExchangeIdleQueue() (*model.Queue, error) {
	eb := ExchangeBusiness{Name: b.Name}
	if b.ExchangeName != "" {
		eb.Name = b.ExchangeName
	}
	if b.ExchangeId != 0 {
		eb.Id = b.ExchangeId
	}

	eId := eb.GetIdByName()
	if eId == 0 {
		return nil, status.Errorf(codes.NotFound, "交换机不存在")
	}

	entity := model.Queue{}
	if res := global.DB.
		Where(model.Queue{ExchangeID: eId}).
		Where("status in ?", []int64{enum.QueueStatusIdle, enum.QueueStatusClose}).
		First(&entity); res.RowsAffected == 0 {
		return b.Create()
	}

	return &entity, nil
}

func (b *QueueBusiness) GetIdByName() int64 {
	entity := model.Queue{}
	res := global.DB.Where(model.Queue{Name: b.Name}).Select("id").First(&entity)
	if res.RowsAffected == 0 {
		return 0
	}
	return entity.ID
}

func (b *QueueBusiness) UpdateByName() error {
	entity := model.Queue{}
	if res := global.DB.Where(model.Queue{Name: b.Name}).First(&entity); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "队列不存在")
	}
	if b.Status != nil {
		entity.Status = *b.Status
	}
	if b.Name != "" {
		entity.Name = b.Name
	}

	global.DB.Save(&entity)
	return nil
}

// PushQueue 发送队列消息
func PushQueue(queueName string, body []byte) {
	ch, _ := global.MessageQueueClient.Channel()
	err := ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			Body: body,
		})
	if err != nil {
		fmt.Printf("send queue message err: %s", err)
	}
}
