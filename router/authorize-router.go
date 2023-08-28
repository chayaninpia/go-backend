package router

import (
	"apps/handler"

	"github.com/gin-gonic/gin"
)

func AuthorizeRouter(e *gin.Engine) {

	authorizeHandler := handler.NewAuthorizeHandler()
	g := e.Group(`api/v1/login`)
	{
		g.POST(``, authorizeHandler.Login)
	}
}
