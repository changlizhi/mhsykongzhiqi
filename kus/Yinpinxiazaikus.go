package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinshijian(moxing *moxings.Yinpinxiazais) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinxiazais----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinshijian(moxing moxings.Yinpinxiazais) *moxings.Yinpinxiazais {
	find := Db.Find(&moxings.Yinpinxiazais{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinshijian--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinxiazais)
	return ret
}
