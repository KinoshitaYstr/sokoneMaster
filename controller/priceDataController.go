package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.mod/service"
)

func AddPriceDataToProductAndStore(c *gin.Context) {
	storeID, err := strconv.Atoi(c.PostForm("store_id"))
	if err != nil {
		panic(err)
	}
	productID, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		panic(err)
	}
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		panic(err)
	}
	service.CreatePriceData(price, productID, storeID)
	c.Redirect(http.StatusFound, "/store/show/"+c.PostForm("store_id"))
}

func UpdatePriceDataByProductPriceByStoreShow(c *gin.Context) {
	priceDataID, err := strconv.Atoi(c.PostForm("price_data_id"))
	if err != nil {
		panic(err)
	}
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		panic(err)
	}
	service.UpdatePriceDataById(priceDataID, price)
	c.Redirect(http.StatusFound, "/store/show/"+c.PostForm("store_id"))
}
