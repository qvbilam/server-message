package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	userProto "message/api/qvbilam/user/v1"
	"message/config"
)

var (
	DB                 *gorm.DB
	Redis              redis.Client
	ServerConfig       *config.ServerConfig
	MessageQueueClient *amqp.Connection
	UserServerClient   userProto.UserClient
)
