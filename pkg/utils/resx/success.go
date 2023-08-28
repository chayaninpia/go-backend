package resx

import (
	"apps/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data any) {

	c.JSON(http.StatusOK, models.ResponseAPI[any]{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}
