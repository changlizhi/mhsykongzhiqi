package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinmimajiu(moxing *moxings.Yinpinmimajius) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinmimajius----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinmimajiu(moxing moxings.Yinpinmimajius) *moxings.Yinpinmimajius {
	find := Jichucaozuo().Find(&moxings.Yinpinmimajius{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinmimajiu--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinmimajius)
	return ret
}
