package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"message/global"
	"message/model/doc"
	"os"
)

func InitElasticSearch() {
	host := global.ServerConfig.ESConfig.Host
	port := global.ServerConfig.ESConfig.Port

	url := elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port))
	sniff := elastic.SetSniff(false) // 不将本地地址转换
	var err error
	logger := log.New(os.Stdout, "elasticsearch", log.LstdFlags) // 设置日志输出位置
	global.ES, err = elastic.NewClient(url, sniff, elastic.SetTraceLog(logger))
	//global.ES, err = elastic.NewClient(url, sniff)
	if err != nil {
		zap.S().Panicf("连接es异常: %s", err.Error())
	}

	// 创建 mapping
	createPrivateChatIndex()
	createGroupChatIndex()
}

func createPrivateChatIndex() {
	exists, err := global.ES.IndexExists(doc.Private{}.GetIndexName()).Do(context.Background())
	if err != nil {
		zap.S().Panicf("私聊索引异常: %s", err)
	}
	if !exists { // 创建索引
		createIndex, err := global.ES.
			CreateIndex(doc.Private{}.GetIndexName()).
			BodyString(doc.Private{}.GetMapping()).
			Do(context.Background())
		if err != nil {
			zap.S().Panicf("创建私聊索引异常: %s", err)
		}

		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func createGroupChatIndex() {
	exists, err := global.ES.IndexExists(doc.Group{}.GetIndexName()).Do(context.Background())
	if err != nil {
		zap.S().Panicf("群聊索引异常: %s", err)
	}
	if !exists { // 创建索引
		createIndex, err := global.ES.
			CreateIndex(doc.Group{}.GetIndexName()).
			BodyString(doc.Group{}.GetMapping()).
			Do(context.Background())
		if err != nil {
			zap.S().Panicf("创建群聊索引异常: %s", err)
		}

		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}
