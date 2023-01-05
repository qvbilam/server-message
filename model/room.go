package model

type Room struct {
	IDModel
	UserModel
	RoomId     int64  `gorm:"type:int not null default 0;comment:房间id;index:idx_room"`
	MessageUid string `gorm:"type:varchar(64) not null;comment:消息uid;index:idx_uid"`
	Type       string `gorm:"type:varchar(64) not null;comment:消息文本内容;"`
	Content    string `gorm:"type:varchar(1024) not null;comment:消息文本内容;"`
	DateModel
	DeletedModel
}
