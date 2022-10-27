package stock

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Read(c *gin.Context) {

	req := &GetStockI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	e, err := xormx.Init()
	if err != nil {
		logx.Error(c, err.Error())
	}

	res, err := req.QueryStock(e)
	if err != nil {
		logx.Error(c, err.Error())
	}

	defer func() {
		if err := e.Close();err != nil{
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, ``, res)

}
func (gs *GetStockI) QueryStock(e *xorm.Engine) ([]GetStockO, error) {

	res := make([]GetStockO,0)

	qs := e.Select(`tp.id, tp.barcode_id, tp.product_name , tps.quantity`).Table(tb.TProductStock{}.TableName()).Alias(`tps`).
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
		res = append(res, GetStockO{
			ProductId:   v[`id`].(string),
			BarcodeId:   v[`barcode_id`].(string),
			ProductName: v[`product_name`].(string),
			Quantity:    int(v[`quantity`].(int32)),
		})
	}

	return res, nil
}
