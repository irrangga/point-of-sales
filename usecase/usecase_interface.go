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
	GetAllProductsUsecase(id string) ([]models.Product, error)
	GetProductUsecase(usr models.Product, id string) (models.Product, error)
	UpdateProductUsecase(usr models.Product, id string) (models.Product, error)
	DeleteProductUsecase(usr models.Product, id string) (models.Product, error)
}

type OutletUsecaseInterface interface {
	AddOutletUsecase(usr models.Outlet) (models.Outlet, error)
	GetAllOutletsUsecase(id string) ([]models.Outlet, error)
	GetOutletUsecase(usr models.Outlet, id string) (models.Outlet, error)
	UpdateOutletUsecase(usr models.Outlet, id string) (models.Outlet, error)
	DeleteOutletUsecase(usr models.Outlet, id string) (models.Outlet, error)
}
