package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	User       User `gorm:"not null"`
	OrderItems []OrderItem
	Status     string `gorm:"type:varchar(100);not null"`
}
