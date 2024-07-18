package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"message/config"
	"message/global"
	"os"
	"strconv"
)

func InitConfig() {
	initEnvConfig()
	initViperConfig()
}

func initEnvConfig() {

	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_HOST"))
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_PASSWORD"))
	ESPort, _ := strconv.Atoi(os.Getenv("ES_PORT"))
	rabbitMQPort, _ := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	userServerPort, _ := strconv.Atoi(os.Getenv("USER_SERVER_PORT"))
	contactServerPort, _ := strconv.Atoi(os.Getenv("CONTACT_SERVER_PORT"))

	if global.ServerConfig == nil {
		global.ServerConfig = &config.ServerConfig{}
	}
	// server
	global.ServerConfig.Name = os.Getenv("SERVER_NAME")
	global.ServerConfig.Port = serverPort
	global.ServerConfig.DBConfig.Host = os.Getenv("DB_HOST")
	// database
	global.ServerConfig.DBConfig.Port = dbPort
	global.ServerConfig.DBConfig.User = os.Getenv("DB_USER")
	global.ServerConfig.DBConfig.Password = os.Getenv("DB_PASSWORD")
	global.ServerConfig.DBConfig.Database = os.Getenv("DB_DATABASE")
	// redis
	global.ServerConfig.RedisConfig.Host = os.Getenv("REDIS_HOST")
	global.ServerConfig.RedisConfig.Port = redisPort
	global.ServerConfig.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
	global.ServerConfig.RedisConfig.Database = redisDb
	// elasticsearch
	global.ServerConfig.ESConfig.Host = os.Getenv("ES_HOST")
	global.ServerConfig.ESConfig.Port = ESPort
	// rabbitmq
	global.ServerConfig.RabbitMQServerConfig.Host = os.Getenv("RABBITMQ_HOST")
	global.ServerConfig.RabbitMQServerConfig.Port = int64(rabbitMQPort)
	global.ServerConfig.RabbitMQServerConfig.Name = os.Getenv("RABBITMQ_NAME")
	global.ServerConfig.RabbitMQServerConfig.User = os.Getenv("RABBITMQ_USER")
	global.ServerConfig.RabbitMQServerConfig.Password = os.Getenv("RABBITMQ_PASSWORD")
	global.ServerConfig.RabbitMQServerConfig.Exchange = os.Getenv("RABBITMQ_EXCHANGE")
	global.ServerConfig.RabbitMQServerConfig.ExchangeChatPrivate = os.Getenv("RABBITMQ_EXCHANGE_CHAT_PRIVATE")
	global.ServerConfig.RabbitMQServerConfig.ExchangeChatGroup = os.Getenv("RABBITMQ_EXCHANGE_CHAT_GROUP")
	global.ServerConfig.RabbitMQServerConfig.ExchangeChatRoom = os.Getenv("RABBITMQ_EXCHANGE_CHAT_ROOM")
	global.ServerConfig.RabbitMQServerConfig.QueuePrefix = os.Getenv("RABBITMQ_QUEUE_SUFFIX")
	// user-server
	global.ServerConfig.UserServerConfig.Name = os.Getenv("USER_SERVER_NAME")
	global.ServerConfig.UserServerConfig.Host = os.Getenv("USER_SERVER_HOST")
	global.ServerConfig.UserServerConfig.Port = int64(userServerPort)
	// contact-server
	global.ServerConfig.ContactServerConfig.Name = os.Getenv("CONTACT_SERVER_NAME")
	global.ServerConfig.ContactServerConfig.Host = os.Getenv("CONTACT_SERVER_HOST")
	global.ServerConfig.ContactServerConfig.Port = int64(contactServerPort)
}

func initViperConfig() {
	file := "config.yaml"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigFile(file)
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}
