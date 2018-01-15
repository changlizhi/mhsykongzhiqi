package moxings

type Yinpinshanchuxins struct {
	Id            int
	Xuliehao      string `gorm:"not null;DEFAULT:0"`
	Yishanchu     int64  `gorm:"not null;DEFAULT:0"`
	Shanchubiaoji int64  `gorm:"not null;DEFAULT:0"`
}

func (Yinpinshanchuxins) TableName() string {
	return "Yinpinshanchuxins"
}
