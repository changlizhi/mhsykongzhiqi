package moxings

type Yinpinshixiaoshijianxins struct {
	Id            int
	Xuliehao      string `gorm:"not null;DEFAULT:0"`
	Shichang int64 `gorm:"not null;DEFAULT:0"`
}

func (Yinpinshixiaoshijianxins) TableName() string {
	return "Yinpinshixiaoshijianxins"
}
