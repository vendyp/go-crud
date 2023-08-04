package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductCategoryId uint
	Name              string `gorm:"size:256"`
	Code              string `gorm:"size:256"`
	Price             int64
}
