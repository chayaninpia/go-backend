package resx

import (
	"apps/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, message string, data any) {

	c.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseAPI[any]{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

func ErrorUnauthorize(c *gin.Context, message string, data any) {

	c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseAPI[any]{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}
