package moxings

import "time"

type Yinpinshijians struct {
	Id              int
	Xuliehao        string    `gorm:"not null;DEFAULT:0"`
	Jieyashijian    time.Time `gorm:"not null;DEFAULT:'1970-01-01 10:00:00'"`
	Dangqianshijian time.Time `gorm:"not null;DEFAULT:'1970-01-01 10:00:00'"`
}

func (Yinpinshijians) TableName() string {
	return "Yinpinshijians"
}
