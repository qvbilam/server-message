package model

import (
	"context"
	"gorm.io/gorm"
	"message/global"
	"message/model/doc"
	"strconv"
)

type Group struct {
	IDModel
	UserModel
	GroupID    int64    `gorm:"type:int not null default 0;comment:群id;index:idx_group"`
	MessageUid string   `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	Type       string   `gorm:"type:varchar(64) not null;comment:消息文本内容;"`
	Content    string   `gorm:"type:varchar(1024) not null;comment:消息文本内容;"`
	Message    *Message `gorm:"foreignKey:uid;references:message_uid"`
	DateModel
	DeletedModel
}

func (entity *Group) AfterCreate(tx *gorm.DB) error {
	d := entity.ToDoc()
	// 写入es
	_, err := global.ES.
		Index().
		Index(doc.Group{}.GetIndexName()).
		BodyJson(d).
		Id(strconv.Itoa(int(entity.ID))).
		Do(context.Background())
	return err
}

func (entity *Group) AfterDelete(tx *gorm.DB) error {
	// 删除 es 数据
	_, err := global.ES.
		Delete().
		Index(doc.Group{}.GetIndexName()).
		Id(strconv.Itoa(int(entity.ID))).
		Do(context.Background())

	return err
}

func (entity *Group) ToDoc() *doc.Group {
	return &doc.Group{
		ID:         entity.ID,
		UserID:     entity.UserID,
		GroupID:    entity.GroupID,
		MessageUID: entity.MessageUid,
		Content:    entity.Content,
	}
}
