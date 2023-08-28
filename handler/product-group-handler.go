package handler

import (
	"apps/models"
	"apps/pkg/productgroup"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type IProductGroupHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Read(c *gin.Context)
}

type ProductGroupHandler struct {
	ProductGroupService productgroup.IProductGroupService
}

func NewProductGroupHandler() *ProductGroupHandler {
	return &ProductGroupHandler{
		ProductGroupService: productgroup.NewProductGroupService(),
	}
}

// ProductGroup godoc
// @Summary      ProductGroupCreate
// @Description  to create product group
// @Tags         ProductGroup
// @Id			 ProductGroupCreate
// @Param   	 CreateProductGroupI  body     models.CreateProductGroupI     false  "CreateProductGroupI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productgroup [post]
func (handler ProductGroupHandler) Create(c *gin.Context) {

	req := models.CreateProductGroupI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductGroupService.Create(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Insert Product Group Success`, nil)
}

// ProductGroup godoc
// @Summary      ProductGroupDelete
// @Description  to delete product group
// @Tags         ProductGroup
// @Id			 ProductGroupDelete
// @Param   	 DeleteProductGroupI  body     models.DeleteProductGroupI     false  "DeleteProductGroupI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productgroup [delete]
func (handler ProductGroupHandler) Delete(c *gin.Context) {

	req := models.DeleteProductGroupI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductGroupService.Delete(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Delete complete`, nil)
}

// ProductGroup godoc
// @Summary      ProductGroupRead
// @Description  to read product group
// @Tags         ProductGroup
// @Id			 ProductGroupRead
// @Param   	 groupName  query     string     false  "groupName"
// @Success      200  {object}  models.ResponseAPI[[]models.ReadProductGroupO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productgroup [get]
func (handler ProductGroupHandler) Read(c *gin.Context) {

	req := models.ReadProductGroupI{}

	req.GroupName = c.Query(`groupName`)

	res, err := handler.ProductGroupService.Read(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// ProductGroup godoc
// @Summary      ProductGroupUpdate
// @Description  to update product group
// @Tags         ProductGroup
// @Id			 ProductGroupUpdate
// @Param   	 UpdateProductGroupI  body     models.UpdateProductGroupI     false  "UpdateProductGroupI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/productgroup [put]
func (handler ProductGroupHandler) Update(c *gin.Context) {

	req := models.UpdateProductGroupI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.ProductGroupService.Update(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, nil)
}
