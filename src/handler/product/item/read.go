package item

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Read(c *gin.Context) {

	req := &ReadProductI{}

	productName := c.Query(`productName`)
	req.ProductName = &productName

	e, err := xormx.Init()
	if err != nil {
		logx.Error(c, err.Error())
	}
	resMain, err := req.Read(e)
	if err != nil {
		logx.Error(c, err.Error())
	}

	res := []ReadProductO{}
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

		itemx := ReadProductO{
			Id:            &id,
			BarcodeId:     &barcodeId,
			ProductName:   &productName,
			ProductNameTh: &productNameTh,
			GroupId:       &groupId,
			BasePrice:     &basePrice,
			SalePrice:     &salePrice,
			IsBaseProduct: &isBaseProduct,
			Unit:          &unit,
		}

		if !v.IsBaseProduct {

			newReq := ReadProductI{ProductId: &v.Id}
			resSub, err := newReq.ReadSub(e)
			if err != nil {
				logx.Error(c, err.Error())
			}

			for _, x := range resSub {
				id := x.Id
				productId := x.ProductId
				baseProductId := x.BaseProductId
				baseProductUse := x.BaseProductUse
				itemx.ProductSub = append(itemx.ProductSub, ReadProductOSub{
					Id:             &id,
					ProductId:      &productId,
					BaseProductId:  &baseProductId,
					BaseProductUse: &baseProductUse,
				})
			}
		}
		res = append(res, itemx)

	}

	resx.Success(c, ``, res)
}

func (rp *ReadProductI) Read(e *xorm.Engine) ([]tb.TProduct, error) {

	res := []tb.TProduct{}

	qs := e.Select(`*`).Table(tb.TProduct{}.TableName())

	if rp.ProductName != nil {
		qs.Where(`product_name LIKE ?`, `%`+*rp.ProductName+`%`)
	}

	if rp.ProductId != nil {
		qs.Where(`product_id = ?`, *rp.ProductId)
	}

	if err := qs.Find(&res); err != nil {
		return nil, err
	}

	return res, nil

}

func (rp *ReadProductI) ReadSub(e *xorm.Engine) ([]tb.TProductSub, error) {

	res := []tb.TProductSub{}

	qs := e.Select(`*`).Table(tb.TProductSub{}.TableName())

	if rp.ProductId != nil {
		qs.Where(`product_id = ?`, *rp.ProductId)
	}

	if err := qs.Find(&res); err != nil {
		return nil, err
	}

	return res, nil
}
