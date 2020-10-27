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
		v1 := storeEngine.Group("/v1")
		{
			v1.POST("/add", controller.AddStore)
			v1.GET("/list", controller.ListStores)
		}
	}

	engine.Run(":3000")
}
