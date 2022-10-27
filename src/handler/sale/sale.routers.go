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
		g.GET(`/channel/read`, channel.Read)
		g.POST(`/channel/create`, channel.Create)
		g.PUT(`/channel/update`, channel.Update)
		g.DELETE(`/channel/delete`, channel.Delete)

		//order
		g.GET(`/order/read`, order.Read)
		g.POST(`/order/create`, order.Create)
	}
}
