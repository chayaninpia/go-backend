package item

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

	req := &CreateProductI{}
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
		if err :=  e.Close();err != nil {
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, `Insert Product Success`, nil)
}

func (cp *CreateProductI) Create(e *xorm.Engine) error {

	newId := uuid.New().String()
	productCreate := tb.TProduct{
		Id:            newId,
		ProductName:   *cp.ProductName,
		ProductNameTh: *cp.ProductNameTh,
		BarcodeId:     *cp.BarcodeId,
		GroupId:       *cp.GroupId,
		BasePrice:     *cp.BasePrice,
		SalePrice:     *cp.SalePrice,
		Unit:          *cp.Unit,
		IsBaseProduct: *cp.IsBaseProduct,
	}

	if _, err := e.Cols(productCreate.Columns()...).Insert(productCreate); err != nil {
		return err
	}

	if _, err := e.Transaction(func(s *xorm.Session) (interface{}, error) {

		if *cp.IsBaseProduct {

			productStockCreate := tb.TProductStock{
				ProductId: newId,
				Quantity:  0,
			}

			if _, err := e.Cols(productStockCreate.Columns()...).Insert(productStockCreate); err != nil {
				if err := s.Rollback(); err != nil{
					return nil, err
				}
				return nil, err
			}
		} else {

			subProductCreate := make([]tb.TProductSub,0)
			for _, v := range cp.SubProduct {

				subProductCreate = append(subProductCreate, tb.TProductSub{
					Id:             uuid.NewString(),
					ProductId:      newId,
					BaseProductId:  *v.BaseProductId,
					BaseProductUse: *v.BaseProductUse,
				})
			}

			if _, err := e.Cols(tb.TProductSub{}.Columns()...).Insert(subProductCreate); err != nil {
				if err := s.Rollback(); err != nil{
					return nil, err
				}
				return nil, err
			}
		}

		return nil, nil
	}); err != nil {

		if _, err := e.Delete(productCreate); err != nil {
			return err
		}
		return err
	}

	return nil
}
