package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:varchar(255);not null"`
	ImageUrl    string  `gorm:"type:varchar(100);not null"`
	Price       float64 `gorm:"not null"`
}
