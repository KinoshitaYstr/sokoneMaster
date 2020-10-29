package model

import (
	"gorm.io/gorm"
)

type PriceData struct {
	gorm.Model
	Price     int
	ProductID int
	Product   Product
	StoreID   int
	Store     Store
}
