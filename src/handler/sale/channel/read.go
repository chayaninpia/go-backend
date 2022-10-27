package channel

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Read(c *gin.Context) {

	req := ReadSaleChannelI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	e, err := xormx.Init()
	if err != nil {
		logx.Error(c, err.Error())
	}

	res, err := req.Read(e)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)

}

func (rsc *ReadSaleChannelI) Read(e *xorm.Engine) ([]ReadSaleChannelO, error) {

	qs := e.Select(`*`).Table(tb.TSaleChannel{}.TableName())

	if rsc.SaleChannel != `` {
		qs.Where(`sale_channel like ?`, `%`+rsc.SaleChannel+`%`)
	}

	result, err := qs.QueryString()
	if err != nil {
		return nil, err
	}

	res := []ReadSaleChannelO{}
	for _, v := range result {
		res = append(res, ReadSaleChannelO{
			Id:          v[`id`],
			SaleChannel: v[`sale_channel`],
		})
	}

	return res, nil
}
