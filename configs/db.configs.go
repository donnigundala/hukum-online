package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDB : initializing mysql database
func Db() *gorm.DB {
	USER := "root"
	PASS := "Asusx42d"
	HOST := "127.0.0.1"
	PORT := "3306"
	DBNAME := "go_crud"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
