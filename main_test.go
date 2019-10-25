package main

import (
	"github.com/jinzhu/gorm"
)

func setupTest() {
	if DB == nil {
		DB, _ = gorm.Open(
			"postgres",
			"host=localhost port=5432 user=coffeeworks dbname=coffeeworks_test password=mypassword sslmode=disable",
		)

		ensureSchema()
	}

	DB.Exec("TRUNCATE reviews")
}
