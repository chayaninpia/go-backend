package order

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/handler/product/item"
	"apps/src/handler/product/stock"
	"apps/src/util/logx"
	"apps/src/util/resx"
	"apps/src/util/validx"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Create(c *gin.Context) {

	req := CreateSaleOrderI{}

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

	res, err := req.Create(e)
	if err != nil {
		logx.Error(c, err.Error())
	}

	defer func() {
		if err := e.Close();err != nil{
			logx.Error(c,err.Error())
		}
	}()

	resx.Success(c, ``, res)
}

func (cso *CreateSaleOrderI) Create(e *xorm.Engine) (*CreateSaleOrderO, error) {

	orderNo := time.Now().Format(`20060102150405`)
	orderDate := time.Now()
	totalBasePrice := 0.0
	totalSalePrice := 0.0
	saleOrderItemCreate := make([]tb.TSaleOrderItem,0)
	updateStock := make([]stock.AddStockI,0)

	for _, v := range cso.Items {

		saleOrderItemCreate = append(saleOrderItemCreate, tb.TSaleOrderItem{
			OrderNo:   orderNo,
			ProductId: v.ProductId,
			Quantity:  v.Quantity,
			SalePrice: v.SalePrice,
			BasePrice: v.BasePrice,
		})

		totalBasePrice += v.BasePrice
		totalSalePrice += v.SalePrice

		//Check Base product
		reqReadSub := item.ReadProductI{
			ProductId: &v.ProductId,
		}

		resReadSub, err := reqReadSub.ReadSub(e)
		if err != nil {
			return nil, err
		}

		if len(resReadSub) == 0 {

			if err := checkStock(e, v.ProductId, v.Quantity, &updateStock); err != nil {
				return nil, err
			}

			continue
		}

		for _, x := range resReadSub {

			totalQuantity := v.Quantity * x.BaseProductUse
			if err := checkStock(e, x.BaseProductId, totalQuantity, &updateStock); err != nil {
				return nil, err
			}
		}

	}

	totalProfit := totalSalePrice - (totalBasePrice + cso.OtherCost)

	saleOrderCreate := tb.TSaleOrder{
		OrderNo:        orderNo,
		SaleType:       cso.SaleType,
		SaleChannelId:  cso.SaleChannelId,
		TotalBasePrice: totalBasePrice,
		TotalSalePrice: totalSalePrice,
		TotalProfit:    totalProfit,
		OtherCost:      cso.OtherCost,
		OrderDate:      orderDate,
	}

	for _, v := range updateStock {
		if err := v.AddStock(e); err != nil {
			return nil, err
		}
	}

	if _, err := e.Cols(saleOrderCreate.Columns()...).Insert(saleOrderCreate); err != nil {
		return nil, err
	}

	if _, err := e.Cols(tb.TSaleOrderItem{}.Columns()...).Insert(saleOrderItemCreate); err != nil {
		return nil, err
	}

	return &CreateSaleOrderO{
		OrderNo:        orderNo,
		SaleType:       cso.SaleType,
		SaleChannelId:  cso.SaleChannelId,
		TotalBasePrice: totalBasePrice,
		TotalSalePrice: totalSalePrice,
		TotalProfit:    totalProfit,
		OtherCost:      cso.OtherCost,
		OrderDate:      orderDate,
	}, nil
}

func checkStock(e *xorm.Engine, productId string, quantity int, updateStock *[]stock.AddStockI) error {

	reqQueryStock := stock.GetStockI{
		ProductId: productId,
	}

	resQueryStock, err := reqQueryStock.QueryStock(e)
	if err != nil {
		return err
	}

	for _, x := range resQueryStock {

		if x.Quantity < quantity {
			return fmt.Errorf(`stock product [ %v ] คงเหลือ [ %v ] ไม่เพียงพอต่อการขาย`, x.ProductName, x.Quantity)
		}

		totalQuantity := x.Quantity - quantity
		*updateStock = append(*updateStock, stock.AddStockI{
			ProductId: productId,
			Quantity:  totalQuantity,
		})
	}

	return nil
}
