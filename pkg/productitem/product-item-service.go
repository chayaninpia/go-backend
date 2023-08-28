package productitem

import (
	"apps/models"
	"apps/pkg/utils/validx"
)

type IProductItemService interface {
	Create(req models.CreateProductI) error
	Delete(req models.DeleteProductI) error
	Read(req models.ReadProductI) ([]models.ReadProductO, error)
	Update(req models.UpdateProductI) error
}

type ProductItemService struct {
	ProductItemRepository IProductItemRepository
}

func NewProductItemService() *ProductItemService {
	return &ProductItemService{
		ProductItemRepository: NewProductItemRepository(),
	}
}

func (service ProductItemService) Create(req models.CreateProductI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	if err := service.ProductItemRepository.Create(req); err != nil {
		return err
	}

	return nil
}

func (service ProductItemService) Delete(req models.DeleteProductI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	if err := service.ProductItemRepository.Delete(req); err != nil {
		return err
	}

	return nil
}

func (service ProductItemService) Read(req models.ReadProductI) ([]models.ReadProductO, error) {

	resMain, err := service.ProductItemRepository.Read(req)
	if err != nil {
		return nil, err
	}

	res := make([]models.ReadProductO, 0)
	for _, v := range resMain {

		id := v.Id
		barcodeId := v.BarcodeId
		productName := v.ProductName
		productNameTh := v.ProductNameTh
		groupId := v.GroupId
		basePrice := v.BasePrice
		salePrice := v.SalePrice
		isBaseProduct := v.IsBaseProduct
		unit := v.Unit

		itemx := models.ReadProductO{
			Id:            id,
			BarcodeId:     barcodeId,
			ProductName:   productName,
			ProductNameTh: productNameTh,
			GroupId:       groupId,
			BasePrice:     basePrice,
			SalePrice:     salePrice,
			IsBaseProduct: isBaseProduct,
			Unit:          unit,
		}

		if !v.IsBaseProduct {

			newReq := models.ReadProductI{ProductId: v.Id}
			resSub, err := service.ProductItemRepository.ReadSub(newReq)
			if err != nil {
				return nil, err
			}

			for _, x := range resSub {
				id := x.Id
				productId := x.ProductId
				baseProductId := x.BaseProductId
				baseProductUse := x.BaseProductUse
				itemx.ProductSub = append(itemx.ProductSub, models.ReadProductOSub{
					Id:             id,
					ProductId:      productId,
					BaseProductId:  baseProductId,
					BaseProductUse: baseProductUse,
				})
			}
		}
		res = append(res, itemx)

	}
	return res, nil
}

func (service ProductItemService) Update(req models.UpdateProductI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}
	if err := service.ProductItemRepository.Update(req); err != nil {
		return err
	}
	return nil
}
