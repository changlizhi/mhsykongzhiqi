package kus

import (
	"log"
	"mhsykongzhiqi/moxings"
	_ "github.com/go-sql-driver/mysql"
)

func Charushebei(moxing *moxings.Shebeis) bool {
	cr := Db.Create(moxing)
	if cr.Error != nil {
		log.Println("Charushebei----cr.Error---", cr.Error)
		return false
	}
	return true
}
