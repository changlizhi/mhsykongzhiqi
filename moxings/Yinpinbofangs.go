package moxings

type Yinpinbofangs struct {
	Id              int
	Xuliehao        string `gorm:"not null;DEFAULT:0"`
	Kaishishijian    int64  `gorm:"not null;DEFAULT:0"`
	Jieshushijian int64  `gorm:"not null;DEFAULT:0"`
}

func (Yinpinbofangs) TableName() string {
	return "Yinpinbofangs"
}
