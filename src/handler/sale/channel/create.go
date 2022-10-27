package channel

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"
	"apps/src/util/validx"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"xorm.io/xorm"
)

func Create(c *gin.Context) {

	req := CreateSaleChannelI{}

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

	if err := req.Create(e); err != nil {
		logx.Error(c, err.Error())
	}

	defer func() {
		if err := e.Close();err != nil{
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, `Add Sale Channel Success`, nil)
}

func (csc *CreateSaleChannelI) Create(e *xorm.Engine) error {

	saleChannelCreate := tb.TSaleChannel{
		Id:          uuid.NewString(),
		SaleChannel: csc.SaleChannel,
	}

	if _, err := e.Cols(saleChannelCreate.Columns()...).Insert(saleChannelCreate); err != nil {
		return nil
	}

	return nil
}
