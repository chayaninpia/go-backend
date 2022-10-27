package order

import (
	"apps/src/db/models/tb"
	"apps/src/db/orm/xormx"
	"apps/src/util/logx"
	"apps/src/util/resx"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func Read(c *gin.Context) {

	req := ReadSaleOrderI{}
	req.OrderNo = c.Query(`orderNo`)

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

func (rso *ReadSaleOrderI) Read(e *xorm.Engine) ([]ReadSaleOrderO, error) {

	qs := e.Select(`*`).Table(tb.TSaleOrder{}.TableName())

	res := make([]ReadSaleOrderO,0)
	resultOrderItems := make([]ReadSaleOrderOItem,0)

	if rso.OrderNo != `` {

		qs.Where(`order_no = ?`, rso.OrderNo)

		resultItems, err := e.Select(`*`).Table(tb.TSaleOrderItem{}.TableName()).
			Where(`order_no = ?`, rso.OrderNo).
			QueryInterface()
		if err != nil {
			return nil, err
		}

		for _, v := range resultItems {
			basePrice, _ := strconv.ParseFloat(string(v[`base_price`].([]uint8)), 64)
			salePrice, _ := strconv.ParseFloat(string(v[`sale_price`].([]uint8)), 64)

			resultOrderItems = append(resultOrderItems, ReadSaleOrderOItem{
				ProductId: v[`product_id`].(string),
				Quantity:  int(v[`quantity`].(int32)),
				BasePrice: basePrice,
				SalePrice: salePrice,
			})
		}

	}

	resultOrder, err := qs.QueryString()
	if err != nil {
		return nil, err
	}

	for _, v := range resultOrder {

		totalBasePrice, _ := strconv.ParseFloat(v[`total_base_price`], 64)
		totalSalePrice, _ := strconv.ParseFloat(v[`total_sale_price`], 64)
		totalProfit, _ := strconv.ParseFloat(v[`total_profit`], 64)
		otherCost, _ := strconv.ParseFloat(v[`other_cost`], 64)
		orderDate, _ := time.Parse(time.RFC3339, v[`order_date`])

		itemx := ReadSaleOrderO{
			OrderNo:        v[`order_no`],
			SaleType:       v[`sale_type`],
			SaleChannelId:  v[`sale_channel_id`],
			TotalBasePrice: totalBasePrice,
			TotalSalePrice: totalSalePrice,
			TotalProfit:    totalProfit,
			OtherCost:      otherCost,
			OrderDate:      orderDate,
		}

		if rso.OrderNo != `` {
			itemx.Items = resultOrderItems
		}

		res = append(res, itemx)
	}

	return res, nil
}
