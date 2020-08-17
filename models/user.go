package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique_index;not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Age      sql.NullInt64
	UserType int64
}
