package kus

import (
	"log"
	"mhsykongzhiqi/moxings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Charuyinpinshijian(moxing *moxings.Yinpinshijians) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Yinpinshijians----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigeyinpinshijian(moxing moxings.Yinpinshijians) *moxings.Yinpinshijians {
	find := Db.Find(&moxings.Yinpinshijians{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigeyinpinshijian--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Yinpinshijians)
	return ret
}
