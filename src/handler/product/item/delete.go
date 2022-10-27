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

func Delete(c *gin.Context) {

	req := DeleteProductI{}

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

	if err := req.Delete(e); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Delete complete`, nil)
}

func (dp *DeleteProductI) Delete(e *xorm.Engine) error {

	productSubDelete := tb.TProductSub{
		ProductId: *dp.Id,
	}
	productDelete := tb.TProduct{
		Id: *dp.Id,
	}

	if _, err := e.Delete(productSubDelete); err != nil {
		return err
	}
	if _, err := e.Delete(productDelete); err != nil {
		return err
	}

	return nil
}
