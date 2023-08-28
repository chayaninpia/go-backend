package router

import (
	"apps/handler"
	"apps/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func MigrateRouter(e *gin.Engine) {

	migrateHandler := handler.NewMigrateHandler()
	g := e.Group(`api/v1/migrate`, middleware.AuthorizationValidate)
	{
		g.POST(`/auto-migrate`, migrateHandler.AutoMigrate)
	}
}
