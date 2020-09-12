package repositories

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DBCon : Database connection
	DBCon *gorm.DB
)

func migrateDatabase() {

	migration, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DATABASE")))

	if err != nil {
		log.Panicf("failed to apply migrations: %s", err.Error())
	}
	if err := migration.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		} else {
			log.Print("No database changes, continue...")
		}
	}
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
		log.Printf("failed to connect database: %s ", err.Error())
		panic(fmt.Sprintf("failed to connect database: %s ", err.Error()))
	}

	migrateDatabase()

}
