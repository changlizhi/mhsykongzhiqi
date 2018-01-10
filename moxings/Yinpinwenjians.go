package moxings

type Yinpinwenjians struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Lujing   string `gorm:"not null;DEFAULT:0"`
	Leixing  string `gorm:"not null;DEFAULT:0"`
	Mima     string `gorm:"not null;DEFAULT:0"`
	Md5mima  string `gorm:"not null;DEFAULT:0"`
}

func (Yinpinwenjians) TableName() string {
	return "Yinpinwenjians"
}
