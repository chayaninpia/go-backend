package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func AddressRouter(e *gin.Engine) {

	addressHandler := handler.NewAddressHandler()
	g := e.Group(`api/v1/address`, middleware.AuthorizationValidate)
	{
		g.POST(``, addressHandler.Create)
		g.GET(``, addressHandler.Read)
		g.PUT(``, addressHandler.Update)
		g.DELETE(``, addressHandler.Delete)
	}
}
