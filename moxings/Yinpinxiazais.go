package moxings

type Yinpinxiazais struct {
	Id              int
	Xuliehao        string `gorm:"not null;DEFAULT:0"`
	Xiazaishijian    int64  `gorm:"not null;DEFAULT:0"`
	Dangqianshijian int64  `gorm:"not null;DEFAULT:0"`
}

func (Yinpinxiazais) TableName() string {
	return "Yinpinxiazais"
}
