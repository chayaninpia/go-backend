package handler

import (
	"apps/models"
	"apps/pkg/saleorder"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type ISaleOrderHandler interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
}

type SaleOrderHandler struct {
	SaleOrderService saleorder.ISaleOrderService
}

func NewSaleOrderHandler() *SaleOrderHandler {
	return &SaleOrderHandler{
		SaleOrderService: saleorder.NewSaleOrderService(),
	}
}

// SaleOrder godoc
// @Summary      SaleOrderCreate
// @Description  to create sale order
// @Tags         SaleOrder
// @Id			 SaleOrderCreate
// @Param   	 SaleOrderCreateI  body     models.SaleOrderCreateI     false  "SaleOrderCreateI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/saleorder [post]
func (handler SaleOrderHandler) Create(c *gin.Context) {

	req := models.SaleOrderCreateI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.SaleOrderService.Create(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Create Order Success`, nil)
}

// SaleOrder godoc
// @Summary      SaleOrderRead
// @Description  to read sale order
// @Tags         SaleOrder
// @Id			 SaleOrderRead
// @Param   	 SaleOrderReadI  body     models.SaleOrderReadI     false  "SaleOrderReadI"
// @Success      200  {object}  models.ResponseAPI[[]models.SaleOrderReadO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/saleorder [get]
func (handler SaleOrderHandler) Read(c *gin.Context) {

	res, err := handler.SaleOrderService.Read()
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// SaleOrder godoc
// @Summary      SaleOrderReadByOrderNo
// @Description  to read sale order by order number
// @Tags         SaleOrder
// @Id			 SaleOrderReadByOrderNo
// @Param   	 orderno  path     string     false  "orderno"
// @Success      200  {object}  models.ResponseAPI[models.SaleOrderReadO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/saleorder/{orderno} [get]
func (handler SaleOrderHandler) ReadByOrderNo(c *gin.Context) {

	orderNo := c.Param("orderno")

	req := models.SaleOrderReadI{}
	req.OrderNo = orderNo

	res, err := handler.SaleOrderService.ReadByOrderNo(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// SaleOrder godoc
// @Summary      SaleOrderUpdate
// @Description  to Update sale order
// @Tags         SaleOrder
// @Id			 SaleOrderUpdate
// @Param   	 SaleOrderUpdateI  body     models.SaleOrderUpdateI     false  "SaleOrderUpdateI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/saleorder/:orderno [get]
func (handler SaleOrderHandler) Update(c *gin.Context) {

	req := models.SaleOrderUpdateI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.SaleOrderService.Update(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `update success`, nil)
}
