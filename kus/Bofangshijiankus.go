package kus

import (
	"log"
	"mhsykongzhiqi/moxings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Charubofangshijian(moxing *moxings.Bofangshijians) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charubofangshijian----cr.Error---", cr.Error)
		return false
	}
	return true
}
func Chaxunyigebofangshijian(moxing moxings.Bofangshijians) *moxings.Bofangshijians {
	find := Db.Find(&moxings.Bofangshijians{}, moxing)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		log.Println("Chaxunyigebofangshijian--find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Bofangshijians)
	return ret
}