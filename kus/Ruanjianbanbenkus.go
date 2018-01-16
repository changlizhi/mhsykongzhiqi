package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuruanjianbanben(moxing *moxings.Ruanjianbanbens) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Ruanjianbanbens----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeruanjianbanben(moxing moxings.Ruanjianbanbens) *moxings.Ruanjianbanbens {
	find := Jichucaozuo().Find(&moxings.Ruanjianbanbens{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeruanjianbanben--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Ruanjianbanbens)
	return ret
}
