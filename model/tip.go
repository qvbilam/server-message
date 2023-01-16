package model

type Tip struct {
	IDModel
	UserModel
	MessageUid string `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	Type       string `gorm:"type:varchar(64) not null;comment:消息文本内容;"`
	Content    string `gorm:"type:varchar(1024) not null;comment:消息文本内容;"`
	DateModel
	DeletedModel
}
