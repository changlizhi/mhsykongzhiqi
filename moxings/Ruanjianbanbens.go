package moxings

type Ruanjianbanbens struct {
	Id       int
	Xuliehao string `gorm:"not null;DEFAULT:0"`
	Banben   string `gorm:"not null;DEFAULT:0"`
}

func (Ruanjianbanbens) TableName() string {
	return "Ruanjianbanbens"
}
