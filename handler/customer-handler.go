package handler

import (
	"apps/models"
	"apps/pkg/customer"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type ICustomerHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Read(c *gin.Context)
}

type CustomerHandler struct {
	CustomerService customer.ICustomerService
}

func NewCustomerHandler() *CustomerHandler {
	return &CustomerHandler{
		CustomerService: customer.NewCustomerService(),
	}
}

// Customer godoc
// @Summary      CustomerCreate
// @Description  to create customer
// @Tags         Customer
// @Id			 CustomerCreate
// @Param   	 CustomerCreateI  body     models.CustomerCreateI     false  "CustomerCreateI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/customer [post]
func (handler CustomerHandler) Create(c *gin.Context) {

	req := models.CustomerCreateI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.CustomerService.Create(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Insert Product Success`, nil)
}

// Customer godoc
// @Summary      CustomerDelete
// @Description  to delete customer
// @Tags         Customer
// @Id			 CustomerDelete
// @Param   	 CustomerDeleteI  body     models.CustomerDeleteI     false  "CustomerDeleteI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/customer [delete]
func (handler CustomerHandler) Delete(c *gin.Context) {

	req := models.CustomerDeleteI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.CustomerService.Delete(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Delete complete`, nil)
}

// Customer godoc
// @Summary      CustomerRead
// @Description  to read customer
// @Tags         Customer
// @Id			 CustomerRead
// @Param   	 CustomerReadI  body     models.CustomerReadI     false  "CustomerReadI"
// @Success      200  {object}  models.ResponseAPI[[]models.ReadProductO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/customer [get]
func (handler CustomerHandler) Read(c *gin.Context) {

	req := models.CustomerReadI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.CustomerService.Read(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// Customer godoc
// @Summary      CustomerUpdate
// @Description  to update customer
// @Tags         Customer
// @Id			 CustomerUpdate
// @Param   	 CustomerUpdateI  body     models.CustomerUpdateI     false  "CustomerUpdateI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/customer [put]
func (handler CustomerHandler) Update(c *gin.Context) {

	req := models.CustomerUpdateI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.CustomerService.Update(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, res)
}
