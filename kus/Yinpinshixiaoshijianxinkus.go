package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuyinpinshixiaoshijianxin(moxing *moxings.Yinpinshixiaoshijianxins) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinshixiaoshijianxins----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinshixiaoshijianxin(moxing moxings.Yinpinshixiaoshijianxins) *moxings.Yinpinshixiaoshijianxins {
	find := Jichucaozuo().Find(&moxings.Yinpinshixiaoshijianxins{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinshixiaoshijianxin--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinshixiaoshijianxins)
	return ret
}
