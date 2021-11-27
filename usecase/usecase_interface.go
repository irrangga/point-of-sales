package usecase

import "point-of-sales/models"

type UserUsecaseInterface interface {
	AddUserUsecase(usr models.Merchant) (models.Merchant, error)
	GetAllUsersUsecase() ([]models.Merchant, error)
	GetUserUsecase(usr models.Merchant, id string) (models.Merchant, error)
	UpdateUserUsecase(usr models.Merchant, id string) (models.Merchant, error)
	DeleteUserUsecase(usr models.Merchant, id string) (models.Merchant, error)
}

type ProductUsecaseInterface interface {
	AddProductUsecase(usr models.Product) (models.Product, error)
	GetAllProductsUsecase() ([]models.Product, error)
	GetProductUsecase(usr models.Product, id string) (models.Product, error)
	UpdateProductUsecase(usr models.Product, id string) (models.Product, error)
	DeleteProductUsecase(usr models.Product, id string) (models.Product, error)
}
