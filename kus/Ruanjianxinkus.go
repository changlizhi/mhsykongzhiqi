package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuruanjianxin(moxing *moxings.Ruanjianxins) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Ruanjianxins----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeruanjianxin(moxing moxings.Ruanjianxins) *moxings.Ruanjianxins {
	find := Jichucaozuo().Find(&moxings.Ruanjianxins{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeruanjianxin--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Ruanjianxins)
	return ret
}