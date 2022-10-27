package migrate

import (
	"apps/src/db/orm/gormx"
	"apps/src/util/logx"
	"apps/src/util/resx"

	"github.com/gin-gonic/gin"
)

func AutoMigrate(c *gin.Context) {

	if err := gormx.AutoMigrate(); err != nil {
		logx.Error(c, err.Error())
	}

	resx.Success(c, `Auto Migrate Complete`, nil)
}
