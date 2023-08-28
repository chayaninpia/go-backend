package handler

import (
	"apps/models"
	"apps/pkg/productstock"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type IProductStockHandler interface {
	Read(c *gin.Context)
	Update(c *gin.Context)
}

type ProductStockHandler struct {
	ProductStockService productstock.IProductStockService
}

func NewProductStockHandler() *ProductStockHandler {
	return &ProductStockHandler{
		ProductStockService: productstock.NewProductStockService(),
	}
}

// ProductStock godoc
// @Summary      ProductStockRead
// @Description  to read product stock
// @Tags         ProductStock
// @Id			 ProductStockRead
// @Param   	 GetStockI  body     models.GetStockI     false  "GetStockI"
// @Success      200  {object}  models.ResponseAPI[[]models.GetStockO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productstock [get]
func (handler ProductStockHandler) Read(c *gin.Context) {

	req := models.GetStockI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.ProductStockService.Read(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// ProductStock godoc
// @Summary      ProductStockUpdate
// @Description  to update product stock
// @Tags         ProductStock
// @Id			 ProductStockUpdate
// @Param   	 AddStockI  body     models.AddStockI     false  "AddStockI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productstock [put]
func (handler ProductStockHandler) Update(c *gin.Context) {

	req := models.AddStockI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.ProductStockService.Update(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, res)
}
