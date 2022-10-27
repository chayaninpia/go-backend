package item

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"
	"apps/src/util/validx"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Update(c *gin.Context) {

	req := UpdateProductI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := validx.Struct(req); err != nil {
		logx.Error(c, err.Error())
	}

	e, err := xormx.Init()
	if err != nil {
		logx.Error(c, err.Error())
	}

	if err := req.Update(e); err != nil {
		logx.Error(c, err.Error())
	}

	defer func() {
		if err := e.Close();err != nil{
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, `Update complete`, nil)
}

func (up *UpdateProductI) Update(e *xorm.Engine) error {

	productUpdate := tb.TProduct{
		ProductName:   *up.ProductName,
		ProductNameTh: *up.ProductNameTh,
		BarcodeId:     *up.BarcodeId,
		GroupId:       *up.GroupId,
		BasePrice:     *up.BasePrice,
		SalePrice:     *up.SalePrice,
		Unit:          *up.Unit,
		IsBaseProduct: *up.IsBaseProduct,
	}

	productUpdateCondi := tb.TProduct{
		Id: *up.Id,
	}

	if _, err := e.Update(productUpdate, productUpdateCondi); err != nil {
		return err
	}

	if len(up.SubProduct) != 0 {
		for _, v := range up.SubProduct {

			subProductUpdate := tb.TProductSub{
				ProductId:      *up.Id,
				BaseProductId:  *v.BaseProductId,
				BaseProductUse: *v.BaseProductUse,
			}
			subProductUpdateCondi := tb.TProductSub{
				Id: *v.Id,
			}

			if _, err := e.Update(subProductUpdate, subProductUpdateCondi); err != nil {
				return err
			}
		}
	}

	return nil
}
