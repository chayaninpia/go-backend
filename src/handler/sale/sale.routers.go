package sale

import (
	"apps/src/handler/sale/channel"
	"apps/src/handler/sale/order"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {

	g := e.Group(`/v1/sale`)
	{
		//channel
		g.GET(`/channel`, channel.Read)
		g.POST(`/channel`, channel.Create)
		g.PUT(`/channel`, channel.Update)
		g.DELETE(`/channel`, channel.Delete)

		//order
		g.GET(`/order`, order.Read)
		g.POST(`/order`, order.Create)
	}
}
