package repository

import "point-of-sales/models"

type UserRepoInterface interface {
	AddUserRepo(usr models.Merchant) (models.Merchant, error)
	GetAllUsersRepo() ([]models.Merchant, error)
	GetUserRepo(usr models.Merchant, id string) (models.Merchant, error)
	UpdateUserRepo(usr models.Merchant, id string) (models.Merchant, error)
	DeleteUserRepo(usr models.Merchant, id string) (models.Merchant, error)
}
