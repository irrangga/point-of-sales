package db

import (
	"fmt"
	"log"
	"point-of-sales/config/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	var configDB = env.Config

	if DB != nil {
		return
	}

	urlDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DBUsername,
		configDB.DBPassword,
		configDB.DBHost,
		configDB.DBPort,
		configDB.DBName,
	)

	DB, err = gorm.Open(mysql.Open(urlDB), &gorm.Config{})
	if err != nil {
		log.Println("cannot connect to DB")
	}

	log.Println("success connect to db")

	AutoMigrate(DB)
}
