package service

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.mod/model"
)

func gormConnect() *gorm.DB {
	host := "0.0.0.0"
	port := "5432"
	userName := "postgres"
	pass := "pass"
	dbName := "test"
	dsn := "host=" + host + " port=" + port + " user=" + userName + " password=" + pass + " dbname=" + dbName + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Migration() {
	db := gormConnect()
	db.AutoMigrate(&model.Store{}, &model.Product{}, &model.PriceData{})
}
