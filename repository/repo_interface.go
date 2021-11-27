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
	AddProductRepo(usr models.Product) (models.Product, error)
	GetAllProductsRepo(id string) ([]models.Product, error)
	GetProductRepo(usr models.Product, id string) (models.Product, error)
	UpdateProductRepo(usr models.Product, id string) (models.Product, error)
	DeleteProductRepo(usr models.Product, id string) (models.Product, error)
}

type OutletRepoInterface interface {
	AddOutletRepo(usr models.Outlet) (models.Outlet, error)
	GetAllOutletsRepo(id string) ([]models.Outlet, error)
	GetOutletRepo(usr models.Outlet, id string) (models.Outlet, error)
	UpdateOutletRepo(usr models.Outlet, id string) (models.Outlet, error)
	DeleteOutletRepo(usr models.Outlet, id string) (models.Outlet, error)
}
