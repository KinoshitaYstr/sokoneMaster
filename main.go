package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main.mod/controller"
	"main.mod/service"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*.tmpl")

	service.Migration()

	storeEngine := engine.Group("/store")
	{
		storeEngine.POST("/add", controller.AddStore)
		storeEngine.GET("/delete/:id", controller.DeleteStore)
		storeEngine.GET("/list", controller.ListStores)
		storeEngine.GET("/show/:id", controller.ShowStore)
	}

	productEngin := engine.Group("/product")
	{
		productEngin.POST("/add", controller.AddProduct)
		productEngin.GET("/delete/:id", controller.DeleteProduct)
		productEngin.GET("/list", controller.ListProducts)
		productEngin.GET("/show/:id", controller.ShowProduct)
		productEngin.POST("/update/list_price", controller.UpdateProductListPrice)
	}

	priceDataEngine := engine.Group("/price_data")
	{
		priceDataEngine.POST("/add", controller.AddPriceDataToProductAndStore)
		priceDataEngine.POST("/store/update", controller.UpdatePriceDataByProductPriceByStoreShow)
	}

	engine.Run(":3000")
}
