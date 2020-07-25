package repositories

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DBCon : Database connection
	DBCon *gorm.DB
)

func InitDatabase() {
	var err error

	DBCon, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		))
	if err != nil {
		panic("failed to connect database")
	}
	defer DBCon.Close()
}
