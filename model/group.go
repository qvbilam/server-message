package model

import "time"

type Group struct {
	IDModel
	UserModel
	GroupID    int64  `gorm:"type:int not null default 0;comment:群id;index:idx_group"`
	MessageUid string `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	IsDelete   bool   `gorm:"type:tinyint(4) not null default 1;comment:是否删除"`
	SentAt     *time.Time
	DateModel
}
