package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinlianjiexin(moxing *moxings.Yinpinlianjiexins) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinlianjiexins----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinlianjiexin(moxing moxings.Yinpinlianjiexins) *moxings.Yinpinlianjiexins {
	find := Jichucaozuo().Find(&moxings.Yinpinlianjiexins{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinlianjiexin--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinlianjiexins)
	return ret
}
