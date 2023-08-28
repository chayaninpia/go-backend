package productstock

import (
	"apps/models"
	tablemodels "apps/models/tablemodels"
	"apps/pkg/utils/xormx"

	"xorm.io/xorm"
)

type IProductStockRepository interface {
	QueryStock(gs models.GetStockI) ([]models.GetStockO, error)
	AddStock(as models.AddStockI) error
}

type ProductStockRepository struct {
	E *xorm.Engine
}

func NewProductStockRepository() *ProductStockRepository {
	return &ProductStockRepository{}
}

func (repo *ProductStockRepository) initDBService() error {
	if repo.E == nil {
		e, err := xormx.Init()
		if err != nil {
			return err
		}
		repo.E = e
	}
	return nil
}

func (repo ProductStockRepository) QueryStock(gs models.GetStockI) ([]models.GetStockO, error) {

	if err := repo.initDBService(); err != nil {
		return nil, err
	}

	res := make([]models.GetStockO, 0)

	qs := repo.E.Select(`tp.id, tp.barcode_id, tp.product_name , tps.quantity`).Table(tablemodels.TProductStock{}.TableName()).Alias(`tps`).
		Join(`INNER`, `t_product AS tp`, `tp.id = tps.product_id`)

	if gs.BarcodeId != `` {
		qs.Where(`tp.barcode_id = ?`, gs.BarcodeId)
	}

	if gs.ProductName != `` {
		qs.Where(`tp.product_name = ?`, gs.ProductName)
	}

	if gs.ProductId != `` {
		qs.Where(`tps.product_id = ?`, gs.ProductId)
	}

	qRes, err := qs.QueryInterface()
	if err != nil {
		return nil, err
	}

	for _, v := range qRes {
		res = append(res, models.GetStockO{
			ProductId:   v[`id`].(string),
			BarcodeId:   v[`barcode_id`].(string),
			ProductName: v[`product_name`].(string),
			Quantity:    int(v[`quantity`].(int32)),
		})
	}

	return res, nil
}

func (repo ProductStockRepository) AddStock(as models.AddStockI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	pdStockCondi := tablemodels.TProductStock{
		ProductId: as.ProductId,
	}
	pdStockUpdate := tablemodels.TProductStock{
		Quantity: as.Quantity,
	}

	if _, err := repo.E.Cols(`quantity`).Update(pdStockUpdate, pdStockCondi); err != nil {
		return err
	}
	return nil
}
