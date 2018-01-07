package moxings

type Shebeis struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Macdizhi string `gorm:"not null;DEFAULT:0"`
	Pici     int    `gorm:"not null;DEFAULT:0"`
}

func (Shebeis) TableName() string {
	return "Shebeis"
}
