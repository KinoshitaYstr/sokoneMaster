package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name    string
	Address string
}
