package salechannel

import (
	"apps/models"
	"apps/pkg/utils/validx"
)

type ISaleChannelService interface {
	Create(req models.CreateSaleChannelI) error
	Delete(req models.DeleteSaleChannelI) error
	Read(req models.ReadSaleChannelI) ([]models.ReadSaleChannelO, error)
	Update(req models.UpdateSaleChannelI) error
}

type SaleChannelService struct {
	SaleChannelRepository ISaleChannelRepository
}

func NewSaleChannelService() *SaleChannelService {
	return &SaleChannelService{
		SaleChannelRepository: NewSaleChannelRepository(),
	}
}

func (service SaleChannelService) Create(req models.CreateSaleChannelI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	if err := service.SaleChannelRepository.Create(req); err != nil {
		return err
	}

	return nil
}

func (service SaleChannelService) Delete(req models.DeleteSaleChannelI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	if err := service.SaleChannelRepository.Delete(req); err != nil {
		return err
	}

	return nil
}

func (service SaleChannelService) Read(req models.ReadSaleChannelI) ([]models.ReadSaleChannelO, error) {

	result, err := service.SaleChannelRepository.Read(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service SaleChannelService) Update(req models.UpdateSaleChannelI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}
	if err := service.SaleChannelRepository.Update(req); err != nil {
		return err
	}
	return nil
}
