package handler

import (
	"apps/models"
	"apps/pkg/address"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type IAddressHandler interface {
	Create(c *gin.Context)
}

type AddressHandler struct {
	AddressService address.IAddressService
}

func NewAddressHandler() *AddressHandler {
	return &AddressHandler{
		AddressService: address.NewAddressService(),
	}
}

// Address godoc
// @Summary      Address
// @Description  to create a new address
// @Tags         Address
// @Id			 Address
// @Param   	 AddressCreateI  body     models.AddressCreateI     false  "AddressCreateI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/address [post]
func (handler AddressHandler) Create(c *gin.Context) {

	file, err := c.FormFile("address")
	if err != nil {
		req := models.AddressCreateI{}

		if err := c.ShouldBindJSON(&req); err != nil {
			logx.Error(c, err.Error())
		}

		if err := handler.AddressService.Create(req); err != nil {
			logx.Error(c, err.Error())
		}
	} else {

		fileOpen, _ := file.Open()
		bytes, err := ioutil.ReadAll(fileOpen)
		if err != nil {
			logx.Error(c, err.Error())
		}

		req := []models.AddressCreateI{}
		if err := json.Unmarshal(bytes, &req); err != nil {
			logx.Error(c, err.Error())
		}

		if err := handler.AddressService.CreateFromJson(req); err != nil {
			logx.Error(c, err.Error())
		}
	}

	resx.Success(c, "Create success", nil)
}

// Address godoc
// @Summary      Address
// @Description  to read address
// @Tags         Address
// @Id			 Address
// @Param   	 AddressReadI  body     models.AddressReadI     false  "AddressReadI"
// @Success      200  {object}  models.ResponseAPI[models.AddressReadO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/address [get]
func (handler AddressHandler) Read(c *gin.Context) {

	req := models.AddressReadI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.AddressService.Read(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, "", res)
}

// Address godoc
// @Summary      Address
// @Description  to update address
// @Tags         Address
// @Id			 Address
// @Param   	 AddressUpdateI  body     models.AddressUpdateI     false  "AddressUpdateI"
// @Success      200  {object}  models.ResponseAPI[models.AddressUpdateO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/address [put]
func (handler AddressHandler) Update(c *gin.Context) {

	req := models.AddressUpdateI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.AddressService.Update(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, "Update success", res)
}

// Address godoc
// @Summary      Address
// @Description  to delete address
// @Tags         Address
// @Id			 Address
// @Param   	 AddressDeleteI  body     models.AddressDeleteI     false  "AddressDeleteI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/address [delete]
func (handler AddressHandler) Delete(c *gin.Context) {

	req := models.AddressDeleteI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.AddressService.Delete(req); err != nil {
		logx.Error(c, err.Error())
	}

	// if err := handler.AddressService.DeleteAll(); err != nil {
	// 	logx.Error(c, err.Error())
	// }

	resx.Success(c, "Delete success", nil)
}
