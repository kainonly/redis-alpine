package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//facade.Register()
	//str, err := cache.GetAdmin()
	//if err != nil {
	//	println(err.Error())
	//}
	//println(str)

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=gorm dbname=postgres password=zt931003")
	if err != nil {
		return
	}

	defer db.Close()
}
