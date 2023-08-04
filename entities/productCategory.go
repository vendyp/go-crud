package entities

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name string `gorm:"size:256"`
}
