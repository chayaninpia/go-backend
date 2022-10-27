package stock

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"
	"apps/src/util/validx"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Update(c *gin.Context) {

	reqAdd := &AddStockI{}
	if err := c.ShouldBindJSON(&reqAdd); err != nil {
		logx.Error(c, err.Error())
	}

	if err := validx.Struct(reqAdd); err != nil {
		logx.Error(c, err.Error())
	}

	e, err := xormx.Init()
	if err != nil {
		logx.Error(c, err.Error())
	}

	reqGet := &GetStockI{
		BarcodeId: reqAdd.BarcodeId,
	}

	resGet, err := reqGet.QueryStock(e)
	if err != nil {
		logx.Error(c, err.Error())
	}

	res := GetStockO{}

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
			resx.Error(c, fmt.Sprintf(`stock product [ %v ] คงเหลือ [ %v ] ไม่เพียงพอต่อการขาย`, v.ProductId, v.Quantity), nil)
			return
		}
	}

	if err := reqAdd.AddStock(e); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

func (as *AddStockI) AddStock(e *xorm.Engine) error {

	pdStockCondi := tb.TProductStock{
		ProductId: as.ProductId,
	}
	pdStockUpdate := tb.TProductStock{
		Quantity: as.Quantity,
	}

	if _, err := e.Cols(`quantity`).Update(pdStockUpdate, pdStockCondi); err != nil {
		return err
	}
	return nil
}
