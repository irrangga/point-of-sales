package usecase

import (
	"point-of-sales/models"
	"point-of-sales/repository"
)

type PriceUsecaseStruct struct {
	prc repository.PriceRepoInterface
}

func NewPriceUsecaseImpl(prc repository.PriceRepoInterface) PriceUsecaseInterface {
	return &PriceUsecaseStruct{prc}
}

func (pu PriceUsecaseStruct) AddPriceUsecase(prc models.Price) (models.Price, error) {
	priceData, err := pu.prc.AddPriceRepo(prc)
	if err != nil {
		return models.Price{}, err
	}
	return priceData, nil
}

func (pu PriceUsecaseStruct) GetAllPricesUsecase(id string) ([]models.Price, error) {
	priceData, err := pu.prc.GetAllPricesRepo(id)
	if err != nil {
		return nil, err
	}
	return priceData, nil
}

func (pu PriceUsecaseStruct) GetPriceUsecase(prc models.Price, id string) (models.Price, error) {
	priceData, err := pu.prc.GetPriceRepo(prc, id)
	if err != nil {
		return models.Price{}, err
	}
	return priceData, nil
}

func (pu PriceUsecaseStruct) UpdatePriceUsecase(prc models.Price, id string) (models.Price, error) {
	priceData, err := pu.prc.UpdatePriceRepo(prc, id)
	if err != nil {
		return models.Price{}, err
	}
	return priceData, nil
}

func (pu PriceUsecaseStruct) DeletePriceUsecase(prc models.Price, id string) (models.Price, error) {
	priceData, err := pu.prc.DeletePriceRepo(prc, id)
	if err != nil {
		return models.Price{}, err
	}
	return priceData, nil
}
