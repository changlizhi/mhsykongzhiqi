package moxings

type Yinpindaxiaoxins struct {
	Id          int
	Xuliehao    string `gorm:"not null;DEFAULT:'0'"`
	Daxiao      string `gorm:"not null;DEFAULT:0"`
	Yasuodaxiao string `gorm:"not null;DEFAULT:0"`
}

func (Yinpindaxiaoxins) TableName() string {
	return "Yinpindaxiaoxins"
}
