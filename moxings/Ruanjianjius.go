package moxings

//用于从控制器获取软件内容，做比对
type Ruanjianjius struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Banben   string `gorm:"not null;DEFAULT:0"`
	Weizhi   string `gorm:"not null;DEFAULT:0"`
	Neirong  string `gorm:"not null;DEFAULT:0"`
	Mingcheng string `gorm:"not null;DEFAULT:0"`
	Leixing  string `gorm:"not null;DEFAULT:0"`
}

func (Ruanjianjius) TableName() string {
	return "Ruanjianjius"
}
