package model

type QueueExchange struct {
	IDModel
	Name   string `gorm:"varchar(64) not null default '';comment:名称"`
	Status int64  `gorm:"type:int not null default 0;comment:状态:-1注销,0等待,1启动"`
	DateModel
}

type Queue struct {
	IDModel
	ExchangeID int64  `gorm:"type:int not null default 0;comment:用户id;index:idx_exchange_id"`
	Name       string `gorm:"varchar(64) not null default '';comment:名称"`
	Status     int64  `gorm:"type:int not null default 0;comment:状态:-1注销,0等待,1启动"`
	DateModel
}
