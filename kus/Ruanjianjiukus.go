package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charuruanjianjiu(moxing *moxings.Ruanjianjius) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Ruanjianjius----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeruanjianjiu(moxing moxings.Ruanjianjius) *moxings.Ruanjianjius {
	find := Db.Find(&moxings.Ruanjianjius{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeruanjianjiu--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Ruanjianjius)
	return ret
}
