package usecase

import (
	"point-of-sales/models"
	"point-of-sales/repository"
)

type ProductUsecaseStruct struct {
	prod repository.ProductRepoInterface
}

func NewProductUsecaseImpl(prod repository.ProductRepoInterface) ProductUsecaseInterface {
	return &ProductUsecaseStruct{prod}
}

func (pu ProductUsecaseStruct) AddProductUsecase(prod models.Product) (models.Product, error) {
	productData, err := pu.prod.AddProductRepo(prod)
	if err != nil {
		return models.Product{}, err
	}
	return productData, nil
}

func (pu ProductUsecaseStruct) GetAllProductsUsecase() ([]models.Product, error) {
	productData, err := pu.prod.GetAllProductsRepo()
	if err != nil {
		return nil, err
	}
	return productData, nil
}

func (pu ProductUsecaseStruct) GetProductUsecase(prod models.Product, id string) (models.Product, error) {
	productData, err := pu.prod.GetProductRepo(prod, id)
	if err != nil {
		return models.Product{}, err
	}
	return productData, nil
}

func (pu ProductUsecaseStruct) UpdateProductUsecase(prod models.Product, id string) (models.Product, error) {
	productData, err := pu.prod.UpdateProductRepo(prod, id)
	if err != nil {
		return models.Product{}, err
	}
	return productData, nil
}

func (pu ProductUsecaseStruct) DeleteProductUsecase(prod models.Product, id string) (models.Product, error) {
	productData, err := pu.prod.DeleteProductRepo(prod, id)
	if err != nil {
		return models.Product{}, err
	}
	return productData, nil
}
