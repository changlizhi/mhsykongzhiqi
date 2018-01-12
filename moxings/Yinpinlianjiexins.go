package moxings

type Yinpinlianjiexins struct {
	Id             int
	Xuliehao       string `gorm:"not null;DEFAULT:0"`
	Ipdizhi        string `gorm:"not null;DEFAULT:0"`
	Duankou        string `gorm:"not null;DEFAULT:0"`
	Lianjie        string `gorm:"not null;DEFAULT:0"`
}

func (Yinpinlianjiexins) TableName() string {
	return "Yinpinlianjiexins"
}
