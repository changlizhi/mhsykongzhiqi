package kus

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mhsykongzhiqi/moxings"
)

func Charushebei(moxing *moxings.Shebeis) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charushebei----cr.Error---", cr.Error)
		return false
	}
	return true
}
