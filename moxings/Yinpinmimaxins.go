package moxings

type Yinpinmimaxins struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Yinpinmima  string `gorm:"not null;DEFAULT:0"`
}

func (Yinpinmimaxins) TableName() string {
	return "Yinpinmimaxins"
}
