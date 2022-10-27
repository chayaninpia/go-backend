package productgroup

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

	req := UpdateProductGroupI{}

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

func (upg *UpdateProductGroupI) Update(e *xorm.Engine) error {

	productGroupUpdate := tb.TProductGroup{
		GroupName: upg.GroupName,
	}
	productGroupUpdateCondi := tb.TProductGroup{
		Id: upg.Id,
	}

	if _, err := e.Update(productGroupUpdate, productGroupUpdateCondi); err != nil {
		return err
	}

	return nil
}
