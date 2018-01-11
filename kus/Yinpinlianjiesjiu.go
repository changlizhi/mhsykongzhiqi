package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinlianjiejiu(moxing *moxings.Yinpinlianjiejius) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinlianjiejius----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinlianjiejiu(moxing moxings.Yinpinlianjiejius) *moxings.Yinpinlianjiejius {
	find := Db.Find(&moxings.Yinpinlianjiejius{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinlianjiejiu--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinlianjiejius)
	return ret
}
