package moxings

import "time"

type Bofangshijians struct {
	Id            int
	Xuliehao      string    `gorm:"not null;DEFAULT:0"`
	Kaishishijian int64  `gorm:"not null;DEFAULT:0"`
	Jieshushijian int64  `gorm:"not null;DEFAULT:0"`
}

func (Bofangshijians) TableName() string {
	return "Bofangshijians"
}
