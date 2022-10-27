package productgroup

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Read(c *gin.Context) {

	req := &ReadProductGroupI{}

	req.GroupName = c.Query(`groupName`)

	e, err := xormx.Init()
	if err != nil {
		logx.Error(c, err.Error())
	}

	result, err := req.Read(e)
	if err != nil {
		logx.Error(c, err.Error())
	}

	res := make([]ReadProductGroupO,0)
	for _, v := range result {

		res = append(res, ReadProductGroupO{
			Id:        v.Id,
			GroupName: v.GroupName,
		})
	}

	defer func() {
		if err := e.Close();err != nil{
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, ``, res)
}

func (rp *ReadProductGroupI) Read(e *xorm.Engine) ([]tb.TProductGroup, error) {

	res := make([]tb.TProductGroup,0)

	qs := e.Select(`*`).Table(tb.TProductGroup{}.TableName())

	if rp.GroupName != `` {
		qs.Where(`group_name LIKE ?`, `%`+rp.GroupName+`%`)
	}

	if err := qs.Find(&res); err != nil {
		return nil, err
	}

	return res, nil
}
