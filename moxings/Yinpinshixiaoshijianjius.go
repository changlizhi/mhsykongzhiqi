package moxings

type Yinpinshixiaoshijianjius struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Shichang int64  `gorm:"not null;DEFAULT:0"`
}

func (Yinpinshixiaoshijianjius) TableName() string {
	return "Yinpinshixiaoshijianjius"
}
