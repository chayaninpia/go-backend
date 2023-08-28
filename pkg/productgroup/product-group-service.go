package productgroup

import (
	"apps/models"
	"apps/pkg/utils/validx"
)

type IProductGroupService interface {
	Create(req models.CreateProductGroupI) error
	Delete(req models.DeleteProductGroupI) error
	Read(req models.ReadProductGroupI) ([]models.ReadProductGroupO, error)
	Update(req models.UpdateProductGroupI) error
}

type ProductGroupService struct {
	ProductGroupRepository IProductGroupRepository
}

func NewProductGroupService() *ProductGroupService {
	return &ProductGroupService{
		ProductGroupRepository: NewProductGroupRepository(),
	}
}

func (service ProductGroupService) Create(req models.CreateProductGroupI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	if err := service.ProductGroupRepository.Create(req); err != nil {
		return err
	}

	return nil
}

func (service ProductGroupService) Delete(req models.DeleteProductGroupI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	if err := service.ProductGroupRepository.Delete(req); err != nil {
		return err
	}

	return nil
}

func (service ProductGroupService) Read(req models.ReadProductGroupI) ([]models.ReadProductGroupO, error) {

	result, err := service.ProductGroupRepository.Read(req)
	if err != nil {
		return nil, err
	}

	res := make([]models.ReadProductGroupO, 0)
	for _, v := range result {

		res = append(res, models.ReadProductGroupO{
			Id:        v.Id,
			GroupName: v.GroupName,
		})
	}
	return res, nil
}

func (service ProductGroupService) Update(req models.UpdateProductGroupI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}
	if err := service.ProductGroupRepository.Update(req); err != nil {
		return err
	}
	return nil
}
