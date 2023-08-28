package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(e *gin.Engine) {

	customerHandler := handler.NewCustomerHandler()
	g := e.Group(`api/v1/customer`, middleware.AuthorizationValidate)
	{
		g.POST(``, customerHandler.Create)
		g.GET(``, customerHandler.Read)
		g.PUT(``, customerHandler.Update)
		g.DELETE(``, customerHandler.Delete)
	}
}
