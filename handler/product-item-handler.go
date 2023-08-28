package handler

import (
	"apps/models"
	"apps/pkg/productitem"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type IProductItemHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Read(c *gin.Context)
}

type ProductItemHandler struct {
	ProductItemService productitem.IProductItemService
}

func NewProductItemHandler() *ProductItemHandler {
	return &ProductItemHandler{
		ProductItemService: productitem.NewProductItemService(),
	}
}

// ProductItem godoc
// @Summary      ProductItemCreate
// @Description  to create product item
// @Tags         ProductItem
// @Id			 ProductItemCreate
// @Param   	 CreateProductI  body     models.CreateProductI     false  "CreateProductI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productitem [post]
func (handler ProductItemHandler) Create(c *gin.Context) {

	req := models.CreateProductI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductItemService.Create(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Insert Product Success`, nil)
}

// ProductItem godoc
// @Summary      ProductItemDelete
// @Description  to delete product item
// @Tags         ProductItem
// @Id			 ProductItemDelete
// @Param   	 DeleteProductI  body     models.DeleteProductI     false  "DeleteProductI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productitem [delete]
func (handler ProductItemHandler) Delete(c *gin.Context) {

	req := models.DeleteProductI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductItemService.Delete(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Delete complete`, nil)
}

// ProductItem godoc
// @Summary      ProductItemRead
// @Description  to read product item
// @Tags         ProductItem
// @Id			 ProductItemRead
// @Param   	 productName  query     string    false  "productName"
// @Success      200  {object}  models.ResponseAPI[[]models.ReadProductO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productitem [get]
func (handler ProductItemHandler) Read(c *gin.Context) {

	req := models.ReadProductI{}

	productName := c.Query(`productName`)
	req.ProductName = productName

	res, err := handler.ProductItemService.Read(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// ProductItem godoc
// @Summary      ProductItemUpdate
// @Description  to update product item
// @Tags         ProductItem
// @Id			 ProductItemUpdate
// @Param   	 UpdateProductI  body     models.UpdateProductI     false  "UpdateProductI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productitem [put]
func (handler ProductItemHandler) Update(c *gin.Context) {

	req := models.UpdateProductI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductItemService.Update(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, nil)
}
