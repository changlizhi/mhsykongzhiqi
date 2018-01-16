package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinmimaxin(moxing *moxings.Yinpinmimaxins) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinmimaxins----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinmimaxin(moxing moxings.Yinpinmimaxins) *moxings.Yinpinmimaxins {
	find := Jichucaozuo().Find(&moxings.Yinpinmimaxins{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinmimaxin--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinmimaxins)
	return ret
}
