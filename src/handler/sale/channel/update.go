package channel

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

	req := UpdateSaleChannelI{}

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

	resx.Success(c, `Update complete`, nil)
}

func (upg *UpdateSaleChannelI) Update(e *xorm.Engine) error {

	saleChannelUpdate := tb.TSaleChannel{
		SaleChannel: upg.SaleChannel,
	}
	saleChannelUpdateCondi := tb.TSaleChannel{
		Id: upg.Id,
	}

	if _, err := e.Update(saleChannelUpdate, saleChannelUpdateCondi); err != nil {
		return err
	}

	return nil
}
