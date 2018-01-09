package moxings

import "time"

type Bofangshijians struct {
	Id            int
	Xuliehao      string    `gorm:"not null;DEFAULT:0"`
	Kaishishijian time.Time `gorm:"not null;DEFAULT:'1970-01-01 10:00:00'"`
	Jieshushijian time.Time `gorm:"not null;DEFAULT:'1970-01-01 10:00:00'"`
}

func (Bofangshijians) TableName() string {
	return "Bofangshijians"
}
