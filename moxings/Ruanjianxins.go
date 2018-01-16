package moxings

//从版本库中拿最新的版本或者拿控制器无法更新到最新版本时找一个合适的版本给控制器
type Ruanjianxins struct {
	Id        int
	Xuliehao  string `gorm:"not null;DEFAULT:0"`
	Banben    string `gorm:"not null;DEFAULT:0"`
	Weizhi    string `gorm:"not null;DEFAULT:0"`
	Mingcheng string `gorm:"not null;DEFAULT:0"`
	Neirong   string `gorm:"not null;DEFAULT:0"`
	Leixing   string `gorm:"not null;DEFAULT:0"`
}

func (Ruanjianxins) TableName() string {
	return "Ruanjianxins"
}
