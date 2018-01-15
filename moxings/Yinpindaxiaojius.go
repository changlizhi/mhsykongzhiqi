package moxings

type Yinpindaxiaojius struct {
	Id          int
	Xuliehao    string `gorm:"not null;DEFAULT:'0'"`
	Daxiao      int    `gorm:"not null;DEFAULT:0"`
	Yasuodaxiao int    `gorm:"not null;DEFAULT:0"`
}

func (Yinpindaxiaojius) TableName() string {
	return "Yinpindaxiaojius"
}
