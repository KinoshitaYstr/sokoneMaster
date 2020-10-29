package service

import "main.mod/model"

func CreateProduct(productName string, listPrice int) {
	db := gormConnect()
	afterDb := db.Create(&model.Product{
		Name:       productName,
		ListPrice:  listPrice,
		PriceDatas: []model.PriceData{},
	})
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}

func DeleteProductById(id int) {
	db := gormConnect()
	target := FindProductById(id)
	afterDb := db.Delete(&target)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}

func FindProductById(id int) model.Product {
	db := gormConnect()
	var product model.Product
	afterDb := db.First(&product, id)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
	return product
}

func FindAllProducts() []model.Product {
	db := gormConnect()
	var products []model.Product
	afterDb := db.Find(&products)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
	return products
}
