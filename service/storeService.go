package service

import (
	"main.mod/model"
)

func CreateStore(storeName string, address string) error {
	db := gormConnect()
	afterDb := db.Create(&model.Store{Name: storeName, Address: address})
	return afterDb.Error
}

func FindStoreById(id int) model.Store {
	db := gormConnect()
	var store model.Store
	db.First(&store, id)
	return store
}

func FindAllStores() []model.Store {
	db := gormConnect()
	var stores []model.Store
	db.Find(&stores)
	return stores
}
