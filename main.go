package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"message/api"
	proto "message/api/qvbilam/message/v1"
	"message/global"
	"message/initialize"
	"message/utils"
	"net"
	"os"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDatabase()
	initialize.InitRedis()
	initialize.InitElasticSearch()
	initialize.InitServer()
	initialize.InitQueue()
	go initialize.InitQueueHealth()

	// 注册服务
	server := grpc.NewServer()
	proto.RegisterMessageServer(server, &api.MessageServer{})

	Host := "0.0.0.0"
	Port, _ := utils.GetFreePort()
	Port = global.ServerConfig.Port

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Host, Port))
	if err != nil {
		zap.S().Panicf("listen port error: %s", err)
	}

	zap.S().Infof("start %s server, host: %s:%d", global.ServerConfig.Name, Host, Port)
	go func() {
		if err := server.Serve(lis); err != nil {
			zap.S().Panicf("start server error: %s", err)
		}
	}()

	// 监听结束
	quit := make(chan os.Signal)
	<-quit
}
