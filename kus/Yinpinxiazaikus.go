package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinxiazai(moxing *moxings.Yinpinxiazais) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinxiazais----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinxiazai(moxing moxings.Yinpinxiazais) *moxings.Yinpinxiazais {
	find := Jichucaozuo().Find(&moxings.Yinpinxiazais{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinxiazai--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinxiazais)
	return ret
}
