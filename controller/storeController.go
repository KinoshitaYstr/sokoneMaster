package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.mod/service"
)

func AddStore(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	if name == "" || address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "parameter error",
		})
	} else {
		err := service.CreateStore(name, address)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusCreated, gin.H{
			"status": "OK",
		})
	}
}

func ListStores(c *gin.Context) {
	storeList := service.FindAllStores()
	c.HTML(http.StatusOK, "store_list.tmpl", gin.H{
		"stores": storeList,
	})
}
