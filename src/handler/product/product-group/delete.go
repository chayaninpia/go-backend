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

func Delete(c *gin.Context) {

	req := DeleteProductGroupI{}

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

func (upg *DeleteProductGroupI) Delete(e *xorm.Engine) error {

	productGroupDelete := tb.TProductGroup{
		Id: upg.Id,
	}

	if _, err := e.Delete(productGroupDelete); err != nil {
		return err
	}

	return nil
}
