package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func SaleChannelRouter(e *gin.Engine) {

	saleChannelHandler := handler.NewSaleChannelHandler()
	g := e.Group(`api/v1/salechannel`, middleware.AuthorizationValidate)
	{
		g.GET(``, saleChannelHandler.Read)
		g.POST(``, saleChannelHandler.Create)
		g.PUT(``, saleChannelHandler.Update)
		g.DELETE(``, saleChannelHandler.Delete)
	}
}
