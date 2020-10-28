package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.mod/service"
)

func AddStore(c *gin.Context) {
	name := c.PostForm("store_name")
	address := c.PostForm("address")
	if name == "" || address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "parameter error: " + name + "-" + address,
		})
	} else {
		err := service.CreateStore(name, address)
		if err != nil {
			panic(err.Error())
		}
		c.Redirect(http.StatusFound, "/store/list")
	}
}

func DeleteStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = service.DeleteStoreById(id)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusFound, "/store/list")
}

func ListStores(c *gin.Context) {
	storeList := service.FindAllStores()
	c.HTML(http.StatusOK, "store_list.tmpl", gin.H{
		"stores": storeList,
	})
}
