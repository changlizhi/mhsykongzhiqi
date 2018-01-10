package kus

import (
	"github.com/jinzhu/gorm"
	"mhsykongzhiqi/moxings"
)

var Db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyshuju?charset=utf8&parseTime=True&loc=Local")

func Chuangjianbiao() {
	Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(
		&moxings.Yinpinwenjians{},
	)
}
