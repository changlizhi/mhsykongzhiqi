package kus

import "github.com/jinzhu/gorm"

var Db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyshuju?charset=utf8&parseTime=True&loc=Local")
