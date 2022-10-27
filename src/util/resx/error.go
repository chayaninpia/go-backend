package resx

import (
	"apps/src/db/models/ut"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, message string, data any) {

	c.JSON(http.StatusInternalServerError, ut.ResX{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}
