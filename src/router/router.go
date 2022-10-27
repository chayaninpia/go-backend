package router

import (
	"apps/src/handler/index"
	"apps/src/handler/migrate"
	"apps/src/handler/product"
	"apps/src/handler/sale"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {

	d := gin.Default()

	product.Router(d)
	migrate.Router(d)
	index.Router(d)
	sale.Router(d)

	hostName, _ := os.Hostname()
	port := `:`
	if strings.Contains(hostName, `local`) {
		port += viper.GetString(`portLocal`)
	} else {
		port += viper.GetString(`port`)
	}

	d.Run(port)
}
