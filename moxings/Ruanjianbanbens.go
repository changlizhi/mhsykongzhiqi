package moxings

//用于记录所有的脚本版本及内容，防止控制器部分更新后找不到正确的版本
type Ruanjianbanbens struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Banben   string `gorm:"not null;DEFAULT:0"`
	Weizhi   string `gorm:"not null;DEFAULT:0"`
	Neirong  string `gorm:"not null;DEFAULT:0"`
}

func (Ruanjianbanbens) TableName() string {
	return "Ruanjianbanbens"
}
