package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func ProductItemRouter(e *gin.Engine) {

	productitemHandler := handler.NewProductItemHandler()
	g := e.Group(`api/v1/productitem`, middleware.AuthorizationValidate)
	{
		g.GET(``, productitemHandler.Read)
		g.POST(``, productitemHandler.Create)
		g.PUT(``, productitemHandler.Update)
		g.DELETE(``, productitemHandler.Delete)
	}
}
