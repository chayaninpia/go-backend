package router

import (
	"apps/pkg/utils/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(d *gin.Engine) {

	d.Use(middleware.CORSMiddleware)

	d.GET("/checkhealth", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "OK") })
	AuthorizeRouter(d)
	MigrateRouter(d)

	// ProductItemRouter(d)
	// ProductGroupRouter(d)
	// ProductStockRouter(d)
	SaleChannelRouter(d)
	SaleOrderRouter(d)
	CustomerRouter(d)
	AddressRouter(d)
	ProductRouter(d)

}
