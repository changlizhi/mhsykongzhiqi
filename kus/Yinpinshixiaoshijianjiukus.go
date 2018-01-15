package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinshixiaoshijianjiu(moxing *moxings.Yinpinshixiaoshijianjius) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinshixiaoshijianjius----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinshixiaoshijianjiu(moxing moxings.Yinpinshixiaoshijianjius) *moxings.Yinpinshixiaoshijianjius {
	find := Db.Find(&moxings.Yinpinshixiaoshijianjius{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinshixiaoshijianjiu--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinshixiaoshijianjius)
	return ret
}
