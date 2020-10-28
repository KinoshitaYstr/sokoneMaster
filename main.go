package main

import (
	"github.com/gin-gonic/gin"
	"main.mod/controller"
	"main.mod/service"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*.tmpl")

	service.Migration()

	storeEngine := engine.Group("/store")
	{
		storeEngine.POST("/add", controller.AddStore)
		storeEngine.GET("/list", controller.ListStores)
	}

	engine.Run(":3000")
}
