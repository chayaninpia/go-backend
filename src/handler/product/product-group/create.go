package productgroup

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"
	"apps/src/util/validx"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"xorm.io/xorm"
)

func Create(c *gin.Context) {

	req := &CreateProductGroupI{}
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

	resx.Success(c, fmt.Sprintf(`Insert Product Group [ %v ] Success`, req.GroupName), nil)
}

func (cpg *CreateProductGroupI) Create(e *xorm.Engine) error {

	pdgInsert := tb.TProductGroup{
		Id:        uuid.New().String(),
		GroupName: cpg.GroupName,
	}

	if _, err := e.Cols(pdgInsert.Columns()...).InsertOne(pdgInsert); err != nil {
		return err
	}

	return nil
}
