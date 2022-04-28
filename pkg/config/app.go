package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"                  //https://v1.gorm.io/docs/
	_ "github.com/jinzhu/gorm/dialects/mysql" //https://go.dev/doc/effective_go#blank
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
