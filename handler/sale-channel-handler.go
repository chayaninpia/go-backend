package handler

import (
	"apps/models"
	"apps/pkg/salechannel"
	"apps/pkg/utils/logx"
	"apps/pkg/utils/resx"

	"github.com/gin-gonic/gin"
)

type ISaleChannelHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Read(c *gin.Context)
}

type SaleChannelHandler struct {
	SaleChannelService salechannel.ISaleChannelService
}

func NewSaleChannelHandler() *SaleChannelHandler {
	return &SaleChannelHandler{
		SaleChannelService: salechannel.NewSaleChannelService(),
	}
}

// SaleChannel godoc
// @Summary      SaleChannelCreate
// @Description  to create sale channel
// @Tags         SaleChannel
// @Id			 SaleChannelCreate
// @Param   	 CreateSaleChannelI  body     models.CreateSaleChannelI     false  "CreateSaleChannelI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/salechannel [post]
func (handler SaleChannelHandler) Create(c *gin.Context) {

	req := models.CreateSaleChannelI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.SaleChannelService.Create(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Insert Sale Channel Success`, nil)
}

// SaleChannel godoc
// @Summary      SaleChannelDelete
// @Description  to delete sale channel
// @Tags         SaleChannel
// @Id			 SaleChannelDelete
// @Param   	 DeleteSaleChannelI  body     models.DeleteSaleChannelI     false  "DeleteSaleChannelI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/salechannel [delete]
func (handler SaleChannelHandler) Delete(c *gin.Context) {

	req := models.DeleteSaleChannelI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.SaleChannelService.Delete(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Delete complete`, nil)
}

// SaleChannel godoc
// @Summary      SaleChannelRead
// @Description  to read sale channel
// @Tags         SaleChannel
// @Id			 SaleChannelRead
// @Param   	 ReadSaleChannelI  body     models.ReadSaleChannelI     false  "ReadSaleChannelI"
// @Success      200  {object}  models.ResponseAPI[[]models.ReadSaleChannelO]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/salechannel [get]
func (handler SaleChannelHandler) Read(c *gin.Context) {

	req := models.ReadSaleChannelI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	res, err := handler.SaleChannelService.Read(req)
	if err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, ``, res)
}

// SaleChannel godoc
// @Summary      SaleChannelUpdate
// @Description  to update sale channel
// @Tags         SaleChannel
// @Id			 SaleChannelUpdate
// @Param   	 UpdateSaleChannelI  body     models.UpdateSaleChannelI     false  "UpdateSaleChannelI"
// @Success      200  {object}  models.ResponseAPI[any]
// @Failure      500  {object}  models.ResponseAPI[any]
// @Router       /api/v1/salechannel [put]
func (handler SaleChannelHandler) Update(c *gin.Context) {

	req := models.UpdateSaleChannelI{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logx.Error(c, err.Error())
	}

	if err := handler.SaleChannelService.Update(req); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Update complete`, nil)
}
