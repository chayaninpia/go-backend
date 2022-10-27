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
		g.GET(`/item/read`, item.Read)
		g.POST(`/item/create`, item.Create)
		g.PUT(`/item/update`, item.Update)
		g.DELETE(`/item/delete`, item.Delete)

		//product-group
		g.GET(`/product-group/read`, productgroup.Read)
		g.POST(`/product-group/create`, productgroup.Create)
		g.PUT(`/product-group/update`, productgroup.Update)
		g.DELETE(`/product-group/delete`, productgroup.Delete)

		//stock
		g.GET(`/stock/read`, stock.Read)
		g.PUT(`/stock/update`, stock.Update)
	}
}
