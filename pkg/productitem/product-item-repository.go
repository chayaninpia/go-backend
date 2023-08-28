package productitem

import (
	"apps/models"
	"apps/models/tablemodels"
	"apps/pkg/utils/xormx"
	"fmt"

	"github.com/google/uuid"
	"xorm.io/xorm"
)

type IProductItemRepository interface {
	Create(cp models.CreateProductI) error
	Delete(dp models.DeleteProductI) error
	Read(rp models.ReadProductI) ([]tablemodels.TProduct, error)
	ReadSub(rp models.ReadProductI) ([]tablemodels.TProductSub, error)
	Update(up models.UpdateProductI) error
}

type ProductItemRepository struct {
	E *xorm.Engine
}

func NewProductItemRepository() *ProductItemRepository {
	return &ProductItemRepository{}
}

func (repo *ProductItemRepository) initDBService() error {
	if repo.E == nil {
		e, err := xormx.Init()
		if err != nil {
			return err
		}
		repo.E = e
	}
	return nil
}

func (repo ProductItemRepository) Create(cp models.CreateProductI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}
	newId := uuid.NewString()
	productCreate := tablemodels.TProduct{
		Id:            newId,
		ProductName:   cp.ProductName,
		ProductNameTh: cp.ProductNameTh,
		BarcodeId:     cp.BarcodeId,
		GroupId:       cp.GroupId,
		BasePrice:     cp.BasePrice,
		SalePrice:     cp.SalePrice,
		Unit:          cp.Unit,
		IsBaseProduct: cp.IsBaseProduct,
	}

	if _, err := repo.E.Cols(productCreate.Columns()...).Insert(productCreate); err != nil {
		return err
	}

	if _, err := repo.E.Transaction(func(s *xorm.Session) (interface{}, error) {

		if cp.IsBaseProduct {

			productStockCreate := tablemodels.TProductStock{
				ProductId: newId,
				Quantity:  0,
			}

			if _, err := s.Cols(productStockCreate.Columns()...).Insert(productStockCreate); err != nil {
				if err := s.Rollback(); err != nil {
					return nil, err
				}
				return nil, err
			}
		} else {

			subProductCreate := make([]tablemodels.TProductSub, 0)
			for _, v := range cp.SubProduct {

				subProductCreate = append(subProductCreate, tablemodels.TProductSub{
					Id:             uuid.NewString(),
					ProductId:      newId,
					BaseProductId:  v.BaseProductId,
					BaseProductUse: v.BaseProductUse,
				})
			}

			if _, err := s.Cols(tablemodels.TProductSub{}.Columns()...).Insert(subProductCreate); err != nil {
				if err := s.Rollback(); err != nil {
					return nil, err
				}
				return nil, err
			}
		}

		return nil, nil
	}); err != nil {

		if _, err := repo.E.Delete(productCreate); err != nil {
			return err
		}
		return err
	}

	return nil
}

func (repo ProductItemRepository) Delete(dp models.DeleteProductI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}
	productSubDelete := tablemodels.TProductSub{
		ProductId: dp.Id,
	}
	productDelete := tablemodels.TProduct{
		Id: dp.Id,
	}

	if _, err := repo.E.Delete(productSubDelete); err != nil {
		return err
	}
	if _, err := repo.E.Delete(productDelete); err != nil {
		return err
	}

	return nil
}

func (repo ProductItemRepository) Read(rp models.ReadProductI) ([]tablemodels.TProduct, error) {

	if err := repo.initDBService(); err != nil {
		return nil, err
	}

	res := make([]tablemodels.TProduct, 0)

	qs := repo.E.Select(`*`).Table(tablemodels.TProduct{}.TableName())

	if rp.ProductName != "" {

		whereLike := fmt.Sprintf("%v%v%v", "%", rp.ProductName, "%")
		qs.Where(`product_name LIKE ?`, whereLike)
	}

	if rp.ProductId != "" {
		qs.Where(`product_id = ?`, rp.ProductId)
	}

	if err := qs.Find(&res); err != nil {
		return nil, err
	}

	return res, nil

}

func (repo ProductItemRepository) ReadSub(rp models.ReadProductI) ([]tablemodels.TProductSub, error) {

	if err := repo.initDBService(); err != nil {
		return nil, err
	}

	res := make([]tablemodels.TProductSub, 0)

	qs := repo.E.Select(`*`).Table(tablemodels.TProductSub{}.TableName())

	if rp.ProductId != "" {
		qs.Where(`product_id = ?`, rp.ProductId)
	}

	if err := qs.Find(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (repo ProductItemRepository) Update(up models.UpdateProductI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	productUpdate := tablemodels.TProduct{
		ProductName:   up.ProductName,
		ProductNameTh: up.ProductNameTh,
		BarcodeId:     up.BarcodeId,
		GroupId:       up.GroupId,
		BasePrice:     up.BasePrice,
		SalePrice:     up.SalePrice,
		Unit:          up.Unit,
		IsBaseProduct: up.IsBaseProduct,
	}

	productUpdateCondi := tablemodels.TProduct{
		Id: up.Id,
	}

	if _, err := repo.E.Update(productUpdate, productUpdateCondi); err != nil {
		return err
	}

	if len(up.SubProduct) != 0 {
		for _, v := range up.SubProduct {

			subProductUpdate := tablemodels.TProductSub{
				ProductId:      up.Id,
				BaseProductId:  v.BaseProductId,
				BaseProductUse: v.BaseProductUse,
			}
			subProductUpdateCondi := tablemodels.TProductSub{
				Id: v.Id,
			}

			if _, err := repo.E.Update(subProductUpdate, subProductUpdateCondi); err != nil {
				return err
			}
		}
	}

	return nil
}
