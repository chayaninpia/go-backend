package resx

import (
	"apps/src/db/models/ut"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data any) {

	c.JSON(http.StatusOK, ut.ResX{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}
