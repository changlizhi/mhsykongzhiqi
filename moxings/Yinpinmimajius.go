package moxings

type Yinpinmimajius struct {
	Id         int
	Xuliehao   string `gorm:"not null;DEFAULT:0"`
	Yinpinmima string `gorm:"not null;DEFAULT:0"`
}

func (Yinpinmimajius) TableName() string {
	return "Yinpinmimajius"
}
