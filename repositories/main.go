package repositories

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/brandomota/golang-api/models"
)

var (
	// DBCon : Database connection
	DBCon *gorm.DB
)

func migrateDatabase() {
	// Add new models on creation
	DBCon.AutoMigrate(models.User{})
}

// InitDatabase : init database connection
func InitDatabase() {
	var err error
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DBCon, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s ", err.Error()))
	}

	migrateDatabase()

	defer DBCon.Close()
}
