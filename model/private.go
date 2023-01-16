package model

import (
	"context"
	"gorm.io/gorm"
	"message/enum"
	"message/global"
	"message/model/doc"
	"strconv"
)

type Private struct {
	IDModel
	UserModel
	TargetUserId int64  `gorm:"type:int not null default 0;comment:接受用户;index:idx_target_user"`
	ChatSn       string `gorm:"type:varchar(64) not null;comment:私聊编号;index:idx_chat_sn"`
	MessageUid   string `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	Type         string `gorm:"type:varchar(64) not null;comment:消息文本内容;"`
	Content      string `gorm:"type:varchar(1024) not null;comment:消息文本内容;"`
	DateModel
	DeletedModel
}

func (entity *Private) AfterCreate(tx *gorm.DB) error {
	var err error
	if entity.Type == enum.MsgTypeTxt {
		d := entity.ToDoc()
		// 写入es
		_, err = global.ES.
			Index().
			Index(doc.Private{}.GetIndexName()).
			BodyJson(d).
			Id(strconv.Itoa(int(entity.ID))).
			Do(context.Background())
	}

	return err
}

func (entity *Private) AfterDelete(tx *gorm.DB) error {
	// 删除 es 数据
	_, err := global.ES.
		Delete().
		Index(doc.Private{}.GetIndexName()).
		Id(strconv.Itoa(int(entity.ID))).
		Do(context.Background())

	return err
}

func (entity *Private) ToDoc() *doc.Private {
	return &doc.Private{
		ID:           entity.ID,
		UserID:       entity.UserID,
		TargetUserID: entity.TargetUserId,
		MessageUID:   entity.MessageUid,
		Content:      entity.Content,
	}
}
