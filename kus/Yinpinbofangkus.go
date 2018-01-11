package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinbofangs(moxing *moxings.Yinpinbofangs) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinbofangs----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinbofangs(moxing moxings.Yinpinbofangs) *moxings.Yinpinbofangs {
	find := Db.Find(&moxings.Yinpinbofangs{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinbofangs--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinbofangs)
	return ret
}
