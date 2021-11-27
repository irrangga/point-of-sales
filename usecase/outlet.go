package usecase

import (
	"point-of-sales/models"
	"point-of-sales/repository"
)

type OutletUsecaseStruct struct {
	out repository.OutletRepoInterface
}

func NewOutletUsecaseImpl(out repository.OutletRepoInterface) OutletUsecaseInterface {
	return &OutletUsecaseStruct{out}
}

func (ou OutletUsecaseStruct) AddOutletUsecase(out models.Outlet) (models.Outlet, error) {
	outletData, err := ou.out.AddOutletRepo(out)
	if err != nil {
		return models.Outlet{}, err
	}
	return outletData, nil
}

func (ou OutletUsecaseStruct) GetAllOutletsUsecase(id string) ([]models.Outlet, error) {
	outletData, err := ou.out.GetAllOutletsRepo(id)
	if err != nil {
		return nil, err
	}
	return outletData, nil
}

func (ou OutletUsecaseStruct) GetOutletUsecase(out models.Outlet, id string) (models.Outlet, error) {
	outletData, err := ou.out.GetOutletRepo(out, id)
	if err != nil {
		return models.Outlet{}, err
	}
	return outletData, nil
}

func (ou OutletUsecaseStruct) UpdateOutletUsecase(out models.Outlet, id string) (models.Outlet, error) {
	outletData, err := ou.out.UpdateOutletRepo(out, id)
	if err != nil {
		return models.Outlet{}, err
	}
	return outletData, nil
}

func (ou OutletUsecaseStruct) DeleteOutletUsecase(out models.Outlet, id string) (models.Outlet, error) {
	outletData, err := ou.out.DeleteOutletRepo(out, id)
	if err != nil {
		return models.Outlet{}, err
	}
	return outletData, nil
}
