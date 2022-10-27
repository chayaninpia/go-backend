package migrate

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {

	g := e.Group(`/v1/migrate`)
	{
		g.GET(`/auto-migrate`, AutoMigrate)
	}
}
