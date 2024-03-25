package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"message/global"
	"message/initialize"
	"message/model"
)

func main() {
	initialize.InitConfig()

	user := global.ServerConfig.DBConfig.User
	password := global.ServerConfig.DBConfig.Password
	host := global.ServerConfig.DBConfig.Host
	port := global.ServerConfig.DBConfig.Port
	database := global.ServerConfig.DBConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
	})
	if err != nil {
		panic(any(err))
	}

	_ = db.AutoMigrate(
		&model.QueueExchange{},
		&model.Queue{},
		&model.Private{},
		&model.Room{},
		&model.Group{},
		&model.Message{},
		&model.System{},
		&model.Tip{},
	)
}
