package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func ProductGroupRouter(e *gin.Engine) {

	productgroupHandler := handler.NewProductGroupHandler()
	g := e.Group(`api/v1/productgroup`, middleware.AuthorizationValidate)
	{
		g.GET(``, productgroupHandler.Read)
		g.POST(``, productgroupHandler.Create)
		g.PUT(``, productgroupHandler.Update)
		g.DELETE(``, productgroupHandler.Delete)
	}
}
