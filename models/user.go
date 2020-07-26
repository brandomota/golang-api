package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  name : string `gorm:"type:varchar(100);unique_index"`
  email: string `gorm:"type:varchar(100);unique_index"`
  age: sql.NullInt64
  type: number
}