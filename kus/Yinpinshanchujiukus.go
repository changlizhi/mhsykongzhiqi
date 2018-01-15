package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Chaxunyigeyinpinshanchujiu(moxing moxings.Yinpinshanchujius) *moxings.Yinpinshanchujius {
	find := Db.Find(&moxings.Yinpinshanchujius{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinshanchujiu--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinshanchujius)
	return ret
}

func Charuyinpinshanchujiu(moxing *moxings.Yinpinshanchujius) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charuyinpinshanchujiu----cr.Error---", cr.Error)
		return false
	}
	return true
}
