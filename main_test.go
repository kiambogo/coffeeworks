package main

import (
	"github.com/kiambogo/coffeeworks/models"
)

func setupTest() {
	if models.DB == nil {
		models.InitializeDB("test")
	}

	models.DB.Exec("TRUNCATE reviews")
	models.DB.Exec("TRUNCATE scores")
}
