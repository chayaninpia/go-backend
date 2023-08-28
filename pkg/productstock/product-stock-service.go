package productstock

import (
	"apps/models"
	"apps/pkg/utils/validx"
	"fmt"
	"strings"
)

type IProductStockService interface {
	Read(req models.GetStockI) ([]models.GetStockO, error)
	Update(reqAdd models.AddStockI) (models.GetStockO, error)
	CheckStock(productId string, quantity int, updateStock []models.AddStockI) error
}

type ProductStockService struct {
	ProductStockRepository IProductStockRepository
}

func NewProductStockService() *ProductStockService {
	return &ProductStockService{
		ProductStockRepository: NewProductStockRepository(),
	}
}

func (service ProductStockService) Read(req models.GetStockI) ([]models.GetStockO, error) {

	res, err := service.ProductStockRepository.QueryStock(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service ProductStockService) Update(reqAdd models.AddStockI) (models.GetStockO, error) {

	if err := validx.Struct(reqAdd); err != nil {
		return models.GetStockO{}, err
	}

	reqGet := models.GetStockI{
		BarcodeId: reqAdd.BarcodeId,
	}

	resGet, err := service.ProductStockRepository.QueryStock(reqGet)
	if err != nil {
		return models.GetStockO{}, err
	}

	res := models.GetStockO{}

	for _, v := range resGet {
		reqAdd.ProductId = v.ProductId
		res.ProductId = v.ProductId
		res.ProductName = v.ProductName
		res.BarcodeId = v.BarcodeId
		if v.BarcodeId == reqAdd.BarcodeId && strings.ToLower(reqAdd.UpdateType) == `add` {
			reqAdd.Quantity += v.Quantity
			res.Quantity = reqAdd.Quantity
		} else if v.BarcodeId == reqAdd.BarcodeId && strings.ToLower(reqAdd.UpdateType) == `del` && v.Quantity > 0 && v.Quantity >= reqAdd.Quantity {
			reqAdd.Quantity = v.Quantity - reqAdd.Quantity
			res.Quantity = reqAdd.Quantity
		} else {
			return models.GetStockO{}, fmt.Errorf(`stock product [ %v ] คงเหลือ [ %v ] ไม่เพียงพอต่อการขาย`, v.ProductId, v.Quantity)
		}
	}

	if err := service.ProductStockRepository.AddStock(reqAdd); err != nil {
		return models.GetStockO{}, err
	}
	return res, nil
}

func (service ProductStockService) CheckStock(productId string, quantity int, updateStock []models.AddStockI) error {

	reqQueryStock := models.GetStockI{
		ProductId: productId,
	}

	resQueryStock, err := service.ProductStockRepository.QueryStock(reqQueryStock)
	if err != nil {
		return err
	}

	for _, x := range resQueryStock {

		if x.Quantity < quantity {
			return fmt.Errorf(`stock product [ %v ] คงเหลือ [ %v ] ไม่เพียงพอต่อการขาย`, x.ProductName, x.Quantity)
		}

		totalQuantity := x.Quantity - quantity
		updateStock = append(updateStock, models.AddStockI{
			ProductId: productId,
			Quantity:  totalQuantity,
		})
	}

	return nil
}
