package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitializeDB(env string) {
	var hostname string
	var err error
	var ok bool

	if hostname, ok = os.LookupEnv("DB_HOST"); !ok {
		hostname = "localhost"
	}

	DB, err = gorm.Open(
		"postgres",
		fmt.Sprintf("host=%v port=5432 user=coffeeworks dbname=coffeeworks_%v password=mypassword sslmode=disable", hostname, env),
	)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err.Error())
	}

	ensureSchema()
}

func ensureSchema() {
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	DB.AutoMigrate(
		&Review{},
		&Score{},
	)
}
