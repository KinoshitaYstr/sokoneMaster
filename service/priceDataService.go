package service

import (
	"main.mod/model"
)

func CreatePriceData(price int, product_id int, store_id int) {
	db := gormConnect()
	product := FindProductById(product_id)
	store := FindStoreById(store_id)
	priceData := model.PriceData{
		Price:   price,
		Product: product,
		Store:   store,
	}
	afterDb := db.Create(&priceData)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}
