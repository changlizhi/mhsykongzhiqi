package moxings

import "time"

type Yinpinshijians struct {
	Id              int
	Xuliehao        string    `gorm:"not null;DEFAULT:0"`
	Jieyashijian    int64  `gorm:"not null;DEFAULT:0"`
	Dangqianshijian int64  `gorm:"not null;DEFAULT:0"`
}

func (Yinpinshijians) TableName() string {
	return "Yinpinshijians"
}
