package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Chaxunyigeyinpinshanchuxin(moxing moxings.Yinpinshanchuxins) *moxings.Yinpinshanchuxins {
	find := Jichucaozuo().Find(&moxings.Yinpinshanchuxins{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinshanchuxin--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinshanchuxins)
	return ret
}

func Charuyinpinshanchuxin(moxing *moxings.Yinpinshanchuxins) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Charuyinpinshanchuxin----cr.Error---", cr.Error)
		return false
	}
	return true
}
