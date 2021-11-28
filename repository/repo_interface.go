package repository

import "point-of-sales/models"

type UserRepoInterface interface {
	AddUserRepo(usr models.Merchant) (models.Merchant, error)
	GetAllUsersRepo() ([]models.Merchant, error)
	GetUserRepo(usr models.Merchant, id string) (models.Merchant, error)
	UpdateUserRepo(usr models.Merchant, id string) (models.Merchant, error)
	DeleteUserRepo(usr models.Merchant, id string) (models.Merchant, error)
}

type ProductRepoInterface interface {
	AddProductRepo(prod models.Product) (models.Product, error)
	GetAllProductsRepo(id string) ([]models.Product, error)
	GetProductRepo(prod models.Product, id string) (models.Product, error)
	UpdateProductRepo(prod models.Product, id string) (models.Product, error)
	DeleteProductRepo(prod models.Product, id string) (models.Product, error)
}

type OutletRepoInterface interface {
	AddOutletRepo(out models.Outlet) (models.Outlet, error)
	GetAllOutletsRepo(id string) ([]models.Outlet, error)
	GetOutletRepo(out models.Outlet, id string) (models.Outlet, error)
	UpdateOutletRepo(out models.Outlet, id string) (models.Outlet, error)
	DeleteOutletRepo(out models.Outlet, id string) (models.Outlet, error)
}

type PriceRepoInterface interface {
	AddPriceRepo(prc models.Price) (models.Price, error)
	GetAllPricesRepo(id string) ([]models.Price, error)
	GetPriceRepo(prc models.Price, id string) (models.Price, error)
	UpdatePriceRepo(prc models.Price, id string) (models.Price, error)
	DeletePriceRepo(prc models.Price, id string) (models.Price, error)
}
