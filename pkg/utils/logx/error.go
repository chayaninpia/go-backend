package logx

import (
	"apps/pkg/utils/resx"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, format string, args ...any) {

	resx.Error(c, fmt.Sprintf(format, args...), nil)
	panic(fmt.Errorf("error : "+format, args...))
}
