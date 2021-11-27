package db

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate()

	if err != nil {
		return
	}
}
