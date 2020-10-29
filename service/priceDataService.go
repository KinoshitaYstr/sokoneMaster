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

func findPriceDataById(id int) model.PriceData {
	var priceData model.PriceData
	db := gormConnect()
	afterDb := db.Find(&priceData, id)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
	return priceData
}

func UpdatePriceDataById(id int, price int) {
	priceData := findPriceDataById(id)
	db := gormConnect()
	priceData.Price = price
	afterDb := db.Save(&priceData)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}

func DeletePriceDataById(id int) {
	db := gormConnect()
	priceData := findPriceDataById(id)
	afterDb := db.Delete(&priceData)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}

type ProductPrice struct {
	PriceDataID uint
	ProductID   uint
	StoreID     uint
	Price       int
	ListPrice   int
	ProductName string
	StoreName   string
}

func FindAllProductPriceByStoreId(storeID int) []ProductPrice {
	db := gormConnect()
	var productPrices []ProductPrice
	var priceDatas []model.PriceData
	store := FindStoreById(storeID)
	db.Where("store_id = ?", storeID).Find(&priceDatas)
	for _, priceData := range priceDatas {
		var productPrice ProductPrice
		productPrice.Price = priceData.Price
		productPrice.PriceDataID = priceData.ID
		productPrice.StoreID = store.ID
		product := FindProductById(priceData.ProductID)
		productPrice.ListPrice = product.ListPrice
		productPrice.ProductID = product.ID
		productPrice.ProductName = product.Name
		productPrice.StoreName = store.Name
		productPrices = append(productPrices, productPrice)
	}
	return productPrices
}

func FindAllProductPriceByProductId(productID int) []ProductPrice {
	db := gormConnect()
	var productPrices []ProductPrice
	var priceDatas []model.PriceData
	product := FindProductById(productID)
	db.Where("product_id = ?", productID).Find(&priceDatas)
	for _, priceData := range priceDatas {
		var productPrice ProductPrice
		productPrice.ProductID = product.ID
		productPrice.Price = priceData.Price
		productPrice.ListPrice = product.ListPrice
		productPrice.PriceDataID = priceData.ID
		store := FindStoreById(priceData.StoreID)
		productPrice.StoreID = store.ID
		productPrice.StoreName = store.Name
		productPrice.ProductName = product.Name
		productPrices = append(productPrices, productPrice)
	}
	return productPrices
}
