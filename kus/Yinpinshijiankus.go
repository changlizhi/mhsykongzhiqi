package kus

import (
	"log"
	"mhsykongzhiqi/moxings"
	_ "github.com/go-sql-driver/mysql"
)

func Charubofangshijian(moxing *moxings.Bofangshijians) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charubofangshijian----cr.Error---", cr.Error)
		return false
	}
	return true
}
