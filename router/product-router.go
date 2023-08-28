package router

import (
	"apps/handler"

	"github.com/gin-gonic/gin"
)

func ProductRouter(e *gin.Engine) {

	productHandler := handler.NewProductHandler()
	g := e.Group(`api/v1/product`)
	{
		g.GET(``, productHandler.Read)
		g.POST(``, productHandler.Create)
		g.PUT(``, productHandler.Update)
		g.DELETE(``, productHandler.Delete)
	}
}
