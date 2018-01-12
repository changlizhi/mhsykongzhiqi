package kus

import (
	"github.com/jinzhu/gorm"
	"github.com/lunny/log"
	"mhsykongzhiqi/moxings"
)

var Db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyshuju?charset=utf8&parseTime=True&loc=Local")

func Chuangjianbiao() {
	db := Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(
		&moxings.Yinpinlianjiejius{},
		&moxings.Yinpinlianjiexins{},
	)
	if db.Error != nil {
		log.Println("Chuangjianbiao---db.Error", db.Error)
	}
}
