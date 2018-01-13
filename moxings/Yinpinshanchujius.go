package moxings

type Yinpinshanchujius struct {
	Id              int
	Xuliehao        string `gorm:"not null;DEFAULT:0"`
	Yishanchu int64  `gorm:"not null;DEFAULT:0"`
	Shanchubiaoji int64  `gorm:"not null;DEFAULT:0"`
}

func (Yinpinshanchujius) TableName() string {
	return "Yinpinshanchujius"
}
