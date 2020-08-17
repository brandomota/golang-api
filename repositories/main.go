package repositories

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

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
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	DBCon, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s ", err.Error()))
	}

	migrateDatabase()

	defer DBCon.Close()
}
