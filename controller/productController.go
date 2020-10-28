package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.mod/service"
)

func AddProduct(c *gin.Context) {
	name := c.PostForm("product_name")
	listPrice, err := strconv.Atoi(c.PostForm("product_list_price"))
	if err != nil {
		panic(err)
	}
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "parameter error: " + name + "-" + c.PostForm("product_list_price"),
		})
	} else {
		service.CreateProduct(name, listPrice)
		c.Redirect(http.StatusFound, "/product/list")
	}
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	service.DeleteProductById(id)
	c.Redirect(http.StatusFound, "/product/list")
}

func ListProducts(c *gin.Context) {
	productList := service.FindAllProducts()
	c.HTML(http.StatusOK, "product_list.tmpl", gin.H{
		"products": productList,
	})
}

func ShowProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	product := service.FindProductById(id)
	c.HTML(http.StatusOK, "product_show.tmpl", gin.H{
		"product": product,
	})
}
