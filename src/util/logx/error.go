package logx

import (
	"apps/src/util/resx"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, format string, args ...any) {

	resx.Error(c, fmt.Sprintf(format, args...), nil)
	panic(fmt.Sprintf("error message : "+format, args...))
}
