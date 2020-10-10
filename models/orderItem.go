package models

type OrderItem struct {
	Order Order     `gorm:"not null"`
	Product Product `gorm:"not null"`
	Quantity int64  `gorm:"not null"`
}
