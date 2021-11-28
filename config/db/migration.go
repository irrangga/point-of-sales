package db

import (
	"point-of-sales/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(
		models.Merchant{},
		models.Outlet{},
		models.Product{},
		models.Price{},
	)

	if err != nil {
		return
	}
}
