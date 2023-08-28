package handler

import (
	"apps/models"
	"apps/pkg/product"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Read(c *gin.Context)
}

type ProductHandler struct {
	ProductService product.IProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		ProductService: product.NewProductService(),
	}
}

// Product godoc
// @Summary      ProductCreate
// @Description  to create product item
// @Tags         Product
// @Id			 ProductCreate
// @Param   	 ProductCreateI  body     models.ProductCreateI     false  "ProductCreateI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/product [post]
func (handler ProductHandler) Create(c *gin.Context) {

	req := models.ProductCreateI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductService.Create(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Insert Product Success`, nil)
}

// Product godoc
// @Summary      ProductDelete
// @Description  to delete product item
// @Tags         Product
// @Id			 ProductDelete
// @Param   	 ProductDeleteI  body     models.ProductDeleteI     false  "ProductDeleteI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/product [delete]
func (handler ProductHandler) Delete(c *gin.Context) {

	req := models.ProductDeleteI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductService.Delete(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Delete complete`, nil)
}

// Product godoc
// @Summary      ProductRead
// @Description  to read product item by id or product name
// @Tags         Product
// @Id			 ProductRead
// @Param   	 productname  query     string    false  "productname"
// @Param   	 id  query     string    false  "id"
// @Success      200  {object}  models.ResponseAPI[[]models.ReadProductO]
// @Success      200  {object}  models.ResponseAPI[models.ReadProductO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/product [get]
func (handler ProductHandler) Read(c *gin.Context) {

	req := models.ProductReadI{}

	if id := c.Query(`id`); len(id) > 0 {
		req.Id = id

		res, err := handler.ProductService.ReadById(req)
		if err != nil {
			logx.Error(c, err.Error())
		}
		resx.Success(c, ``, res)
	} else if productName := c.Query(`productname`); len(productName) > 0 {
		productName := c.Query(`productname`)
		req.NameEn = productName
		req.NameTh = productName

		res, err := handler.ProductService.Read(req)
		if err != nil {
			logx.Error(c, err.Error())
		}
		resx.Success(c, ``, res)
	} else {
		res, err := handler.ProductService.ReadAll()
		if err != nil {
			logx.Error(c, err.Error())
		}
		resx.Success(c, ``, res)
	}

}

// Product godoc
// @Summary      ProductUpdate
// @Description  to update product item
// @Tags         Product
// @Id			 ProductUpdate
// @Param   	 ProductUpdateI  body     models.ProductUpdateI     false  "ProductUpdateI"
// @Success      200  {object}  models.ResponseAPI[models.UpdateProductO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/product [put]
func (handler ProductHandler) Update(c *gin.Context) {

	req := models.ProductUpdateI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.ProductService.Update(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, res)
}

// Product godoc
// @Summary      ProductUpdate
// @Description  to update product item
// @Tags         Product
// @Id			 ProductUpdate
// @Param   	 ProductUpdateI  body     models.ProductUpdateI     false  "ProductUpdateI"
// @Success      200  {object}  models.ResponseAPI[models.UpdateProductO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/product [put]
func (handler ProductHandler) UpdateStock(c *gin.Context) {

	req := models.ProductUpdateI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.ProductService.Update(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, res)
}
