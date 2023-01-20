package product

import (
	"apps/src/handler/product/item"
	productgroup "apps/src/handler/product/product-group"
	"apps/src/handler/product/stock"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {

	g := e.Group(`/v1/product`)
	{
		//item
		g.GET(`/item`, item.Read)
		g.POST(`/item`, item.Create)
		g.PUT(`/item`, item.Update)
		g.DELETE(`/item`, item.Delete)

		//product-group
		g.GET(`/product-group`, productgroup.Read)
		g.POST(`/product-group`, productgroup.Create)
		g.PUT(`/product-group`, productgroup.Update)
		g.DELETE(`/product-group`, productgroup.Delete)

		//stock
		g.GET(`/stock`, stock.Read)
		g.PUT(`/stock`, stock.Update)
	}
}
