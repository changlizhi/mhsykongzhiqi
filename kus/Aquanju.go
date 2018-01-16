package kus

import (
	"github.com/jinzhu/gorm"
	"github.com/lunny/log"
	"mhsykongzhiqi/moxings"
)

func Jichucaozuo() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyshuju?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("链接数据库失败")
		return nil
	}
	return db
}

func Chuangjianbiao() {
	db := Jichucaozuo().Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(
		&moxings.Yinpinlianjiejius{},
		&moxings.Yinpinlianjiexins{},
	)
	if db.Error != nil {
		log.Println("Chuangjianbiao---db.Error", db.Error)
	}
}
