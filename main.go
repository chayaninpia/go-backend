package main

import (
	"apps/src/router"
	"apps/src/util/configx"
)

func main() {

	configx.Init()
	router.Init()
}
