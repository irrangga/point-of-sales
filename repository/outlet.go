package repository

import (
	"point-of-sales/models"

	"gorm.io/gorm"
)

type OutletRepoStruct struct {
	db *gorm.DB
}

func NewOutletRepoImpl(db *gorm.DB) OutletRepoInterface {
	return &OutletRepoStruct{db}
}

func (pr OutletRepoStruct) AddOutletRepo(out models.Outlet) (models.Outlet, error) {
	err := pr.db.Debug().Create(&out).Error
	if err != nil {
		return models.Outlet{}, err
	}
	return out, nil
}

func (pr OutletRepoStruct) GetAllOutletsRepo(id string) ([]models.Outlet, error) {
	var out []models.Outlet
	err := pr.db.Debug().Where("merchant_id = ?", id).Find(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (pr OutletRepoStruct) GetOutletRepo(out models.Outlet, id string) (models.Outlet, error) {
	err := pr.db.Debug().Where("id = ?", id).First(&out).Error
	if err != nil {
		return models.Outlet{}, err
	}
	return out, nil
}

func (pr OutletRepoStruct) UpdateOutletRepo(out models.Outlet, id string) (models.Outlet, error) {
	err := pr.db.Debug().Model(&out).Where("id = ?", id).Updates(&out).Error
	if err != nil {
		return models.Outlet{}, err
	}
	return out, nil
}

func (pr OutletRepoStruct) DeleteOutletRepo(out models.Outlet, id string) (models.Outlet, error) {
	err := pr.db.Debug().Where("id = ?", id).Delete(&out).Error
	if err != nil {
		return models.Outlet{}, err
	}
	return out, nil
}
