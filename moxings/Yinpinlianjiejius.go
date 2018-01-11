package moxings

type Yinpinlianjiejius struct {
	Id              int
	Xuliehao        string `gorm:"not null;DEFAULT:0"`
	Lianjie        string `gorm:"not null;DEFAULT:0"`
}

func (Yinpinlianjiejius) TableName() string {
	return "Yinpinlianjiejius"
}
