package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func ProductStockRouter(e *gin.Engine) {

	productstockHandler := handler.NewProductStockHandler()
	g := e.Group(`api/v1/productstock`, middleware.AuthorizationValidate)
	{
		g.GET(``, productstockHandler.Read)
		g.PUT(``, productstockHandler.Update)
	}
}
