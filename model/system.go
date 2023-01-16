package model

type System struct {
	IDModel
	UserID     int64  `gorm:"type:int not null default 0;comment:用户id;index:idx_user_object"`
	Object     string `gorm:"type:varchar(64) not null;comment:对象;index:idx_user_object"`
	MessageUid string `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	Type       string `gorm:"type:varchar(64) not null;comment:消息文本内容;"`
	Content    string `gorm:"type:varchar(1024) not null;comment:消息文本内容;"`
	DateModel
	DeletedModel
}
