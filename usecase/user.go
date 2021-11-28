package usecase

import (
	"point-of-sales/models"
	"point-of-sales/repository"
)

type UserUsecaseStruct struct {
	usr repository.UserRepoInterface
}

func NewUserUsecaseImpl(usr repository.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecaseStruct{usr}
}

func (uu UserUsecaseStruct) AddUserUsecase(usr models.Merchant) (models.Merchant, error) {
	userData, err := uu.usr.AddUserRepo(usr)
	if err != nil {
		return models.Merchant{}, err
	}
	return userData, nil
}

func (uu UserUsecaseStruct) GetAllUsersUsecase() ([]models.Merchant, error) {
	userData, err := uu.usr.GetAllUsersRepo()
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (uu UserUsecaseStruct) GetUserUsecase(usr models.Merchant, id string) (models.Merchant, error) {
	userData, err := uu.usr.GetUserRepo(usr, id)
	if err != nil {
		return models.Merchant{}, err
	}
	return userData, nil
}

func (uu UserUsecaseStruct) UpdateUserUsecase(usr models.Merchant, id string) (models.Merchant, error) {
	userData, err := uu.usr.UpdateUserRepo(usr, id)
	if err != nil {
		return models.Merchant{}, err
	}
	return userData, nil
}

func (uu UserUsecaseStruct) DeleteUserUsecase(usr models.Merchant, id string) (models.Merchant, error) {
	userData, err := uu.usr.DeleteUserRepo(usr, id)
	if err != nil {
		return models.Merchant{}, err
	}
	return userData, nil
}
