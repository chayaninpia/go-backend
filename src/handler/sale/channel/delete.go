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

func Delete(c *gin.Context) {

	req := DeleteSaleChannelI{}

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

	defer func() {
		if err := e.Close();err != nil{
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, `Delete complete`, nil)
}

func (upg *DeleteSaleChannelI) Delete(e *xorm.Engine) error {

	saleChannelDelete := tb.TSaleChannel{
		Id: upg.Id,
	}

	if _, err := e.Delete(saleChannelDelete); err != nil {
		return err
	}

	return nil
}
