package controller

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.mod/service"

	"googlemaps.github.io/maps"
)

func AddStore(c *gin.Context) {
	name := c.PostForm("store_name")
	address := c.PostForm("address")
	if name == "" || address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "parameter error: " + name + "-" + address,
		})
	} else {
		client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAP_API_KEY")))
		if err != nil {
			panic(err)
		}
		req := &maps.GeocodingRequest{
			Address:  address,
			Language: "ja",
			Region:   "JP",
		}
		res, err := client.Geocode(context.Background(), req)
		if err != nil {
			panic(err)
		}
		service.CreateStore(name, res[0].FormattedAddress, res[0].Geometry.Location.Lat, res[0].Geometry.Location.Lng)
		c.Redirect(http.StatusFound, "/store/list")
	}
}

func DeleteStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	service.DeleteStoreById(id)
	c.Redirect(http.StatusFound, "/store/list")
}

func ListStores(c *gin.Context) {
	storeList := service.FindAllStores()
	c.HTML(http.StatusOK, "store_list.tmpl", gin.H{
		"stores":  storeList,
		"API_KEY": os.Getenv("GOOGLE_MAP_API_KEY"),
	})
}

func ShowStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	store := service.FindStoreById(id)
	// c.Redirect(http.StatusFound, "/store/list")
	c.HTML(http.StatusOK, "store_show.tmpl", gin.H{
		"store":   store,
		"API_KEY": os.Getenv("GOOGLE_MAP_API_KEY"),
	})
}
