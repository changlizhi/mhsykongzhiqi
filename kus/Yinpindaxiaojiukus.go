package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpindaxiaojiu(moxing *moxings.Yinpindaxiaojius) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpindaxiaojius----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpindaxiaojiu(moxing moxings.Yinpindaxiaojius) *moxings.Yinpindaxiaojius {
	find := Db.Find(&moxings.Yinpindaxiaojius{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpindaxiaojiu--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpindaxiaojius)
	return ret
}
