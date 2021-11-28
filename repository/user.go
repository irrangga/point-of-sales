package repository

import (
	"point-of-sales/models"

	"gorm.io/gorm"
)

type UserRepoStruct struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserRepoStruct{db}
}

func (ur UserRepoStruct) AddUserRepo(usr models.Merchant) (models.Merchant, error) {
	err := ur.db.Debug().Create(&usr).Error
	if err != nil {
		return models.Merchant{}, err
	}
	return usr, nil
}

func (ur UserRepoStruct) GetAllUsersRepo() ([]models.Merchant, error) {
	var usr []models.Merchant
	err := ur.db.Debug().Find(&usr).Error
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (ur UserRepoStruct) GetUserRepo(usr models.Merchant, id string) (models.Merchant, error) {
	err := ur.db.Debug().Where("id = ?", id).First(&usr).Error
	if err != nil {
		return models.Merchant{}, err
	}
	return usr, nil
}

func (ur UserRepoStruct) UpdateUserRepo(usr models.Merchant, id string) (models.Merchant, error) {
	err := ur.db.Debug().Model(&usr).Where("id = ?", id).Updates(&usr).Error
	if err != nil {
		return models.Merchant{}, err
	}
	return usr, nil
}

func (ur UserRepoStruct) DeleteUserRepo(usr models.Merchant, id string) (models.Merchant, error) {
	err := ur.db.Debug().Where("id = ?", id).Delete(&usr).Error
	if err != nil {
		return models.Merchant{}, err
	}
	return usr, nil
}
