package service

import (
	"main.mod/model"
)

func CreateStore(storeName string, address string, lat float64, lng float64) {
	db := gormConnect()
	afterDb := db.Create(&model.Store{
		Name:       storeName,
		Address:    address,
		Lat:        lat,
		Lng:        lng,
		PriceDatas: []model.PriceData{},
	})
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}

func DeleteStoreById(id int) {
	db := gormConnect()
	target := FindStoreById(id)
	afterDb := db.Delete(&target)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
}

func FindStoreById(id int) model.Store {
	db := gormConnect()
	var store model.Store
	afterDb := db.First(&store, id).Association("PriceDatas")
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}

	var pd model.PriceData
	var tmp model.Store
	db.First(&pd, 1).Association("store").Find(&tmp)
	return store
}

func FindAllStores() []model.Store {
	db := gormConnect()
	var stores []model.Store
	afterDb := db.Find(&stores)
	if afterDb.Error != nil {
		panic(afterDb.Error)
	}
	return stores
}
