package kus

import (
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
	_ "github.com/go-sql-driver/mysql"
)

var Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyshuju?charset=utf8&parseTime=True&loc=Local")
func Charushebei(moxing *moxings.Shebeis) bool {
	log.Println(err)
	log.Println("moxing-------",moxing)
	log.Println("Db---",Db)
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charushebei----cr.Error---", cr.Error)
		return false
	}
	return true
}
