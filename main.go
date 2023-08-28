// @title           Go Backend POS API
// @version         1.0
// @description     This is a docs of Go Backend POS API

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
// @description Example: "Bearer {token}"
package main

import (
	"apps/pkg/utils/configx"
	"apps/router"
	"fmt"

	_ "apps/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ginE := gin.Default()
	configx.Init()
	router.Init(ginE)

	port := configx.GetPort()

	ginE.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("[SWAGGER] PATH : http://localhost%v/swagger/index.html\n", port)

	ginE.Run(port)
}
