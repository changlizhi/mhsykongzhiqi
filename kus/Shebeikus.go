package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mhsykongzhiqi/moxings"
)

func Chaxunyigeshebei(moxing moxings.Shebeis) *moxings.Shebeis {
	find := Jichucaozuo().Find(&moxings.Shebeis{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeshebei--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Shebeis)
	return ret
}

func Charushebei(moxing *moxings.Shebeis) bool {
	cr := Jichucaozuo().Create(moxing)
	if cr.Error != nil {
		log.Println("Charushebei----cr.Error---", cr.Error)
		return false
	}
	return true
}
