package model

type Message struct {
	IDModel
	Uid     string `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	Type    string `gorm:"type:varchar(64) not null;comment:消息类型"`
	Content string `gorm:"type:text;comment:消息内容"`
	DateModel
}
