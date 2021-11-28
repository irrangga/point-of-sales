package repository

import (
	"point-of-sales/models"

	"gorm.io/gorm"
)

type PriceRepoStruct struct {
	db *gorm.DB
}

func NewPriceRepoImpl(db *gorm.DB) PriceRepoInterface {
	return &PriceRepoStruct{db}
}

func (pr PriceRepoStruct) AddPriceRepo(prc models.Price) (models.Price, error) {
	err := pr.db.Debug().Create(&prc).Error
	if err != nil {
		return models.Price{}, err
	}
	return prc, nil
}

func (pr PriceRepoStruct) GetAllPricesRepo(id string) ([]models.Price, error) {
	var prc []models.Price
	err := pr.db.Debug().Where("product_id = ?", id).Find(&prc).Error
	if err != nil {
		return nil, err
	}
	return prc, nil
}

func (pr PriceRepoStruct) GetPriceRepo(prc models.Price, id string) (models.Price, error) {
	err := pr.db.Debug().Where("id = ?", id).First(&prc).Error
	if err != nil {
		return models.Price{}, err
	}
	return prc, nil
}

func (pr PriceRepoStruct) UpdatePriceRepo(prc models.Price, id string) (models.Price, error) {
	err := pr.db.Debug().Model(&prc).Where("id = ?", id).Updates(&prc).Error
	if err != nil {
		return models.Price{}, err
	}
	return prc, nil
}

func (pr PriceRepoStruct) DeletePriceRepo(prc models.Price, id string) (models.Price, error) {
	err := pr.db.Debug().Where("id = ?", id).Delete(&prc).Error
	if err != nil {
		return models.Price{}, err
	}
	return prc, nil
}
