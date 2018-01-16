package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpindaxiaoxin(moxing *moxings.Yinpindaxiaoxins) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpindaxiaoxins----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpindaxiaoxin(moxing moxings.Yinpindaxiaoxins) *moxings.Yinpindaxiaoxins {
	find := Jichucaozuo().Find(&moxings.Yinpindaxiaoxins{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpindaxiaoxin--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpindaxiaoxins)
	return ret
}
