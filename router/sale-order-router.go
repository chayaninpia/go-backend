package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func SaleOrderRouter(e *gin.Engine) {

	saleOrderHandler := handler.NewSaleOrderHandler()
	g := e.Group(`api/v1/saleorder`, middleware.AuthorizationValidate)
	{
		g.POST(``, saleOrderHandler.Create)
		g.GET(``, saleOrderHandler.Read)
		g.GET(`/:orderno`, saleOrderHandler.ReadByOrderNo)
		g.PUT(``, saleOrderHandler.Update)
	}
}
