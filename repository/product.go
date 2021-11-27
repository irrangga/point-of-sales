package repository

import (
	"point-of-sales/models"

	"gorm.io/gorm"
)

type ProductRepoStruct struct {
	db *gorm.DB
}

func NewProductRepoImpl(db *gorm.DB) ProductRepoInterface {
	return &ProductRepoStruct{db}
}

func (pr ProductRepoStruct) AddProductRepo(prod models.Product) (models.Product, error) {
	err := pr.db.Debug().Create(&prod).Error
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (pr ProductRepoStruct) GetAllProductsRepo() ([]models.Product, error) {
	var prod []models.Product
	err := pr.db.Debug().Find(&prod).Error
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (pr ProductRepoStruct) GetProductRepo(prod models.Product, id string) (models.Product, error) {
	err := pr.db.Debug().Where("id = ?", id).First(&prod).Error
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (pr ProductRepoStruct) UpdateProductRepo(prod models.Product, id string) (models.Product, error) {
	err := pr.db.Debug().Model(&prod).Where("id = ?", id).Updates(&prod).Error
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (pr ProductRepoStruct) DeleteProductRepo(prod models.Product, id string) (models.Product, error) {
	err := pr.db.Debug().Where("id = ?", id).Delete(&prod).Error
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}
