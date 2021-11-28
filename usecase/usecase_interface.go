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
	AddProductUsecase(prod models.Product) (models.Product, error)
	GetAllProductsUsecase(id string) ([]models.Product, error)
	GetProductUsecase(prod models.Product, id string) (models.Product, error)
	UpdateProductUsecase(prod models.Product, id string) (models.Product, error)
	DeleteProductUsecase(prod models.Product, id string) (models.Product, error)
}

type OutletUsecaseInterface interface {
	AddOutletUsecase(out models.Outlet) (models.Outlet, error)
	GetAllOutletsUsecase(id string) ([]models.Outlet, error)
	GetOutletUsecase(out models.Outlet, id string) (models.Outlet, error)
	UpdateOutletUsecase(out models.Outlet, id string) (models.Outlet, error)
	DeleteOutletUsecase(out models.Outlet, id string) (models.Outlet, error)
}

type PriceUsecaseInterface interface {
	AddPriceUsecase(prc models.Price) (models.Price, error)
	GetAllPricesUsecase(id string) ([]models.Price, error)
	GetPriceUsecase(prc models.Price, id string) (models.Price, error)
	UpdatePriceUsecase(prc models.Price, id string) (models.Price, error)
	DeletePriceUsecase(prc models.Price, id string) (models.Price, error)
}
