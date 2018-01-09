package kus

import (
	"log"
	"mhsykongzhiqi/moxings"
	_ "github.com/go-sql-driver/mysql"
)

func Charuyinpinshijian(moxing *moxings.Yinpinshijians) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charuyinpinshijian----cr.Error---", cr.Error)
		return false
	}
	return true
}
