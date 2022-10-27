package index

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {

	g := e.Group(`/v1/index`)
	{
		g.GET(`/get`, Get)
	}
}
