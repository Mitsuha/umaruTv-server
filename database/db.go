package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"umarutv/config"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(config.DB.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}
