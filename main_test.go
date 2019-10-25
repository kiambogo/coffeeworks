package main

import (
	"github.com/kiambogo/coffeeworks/clients"
	"github.com/kiambogo/coffeeworks/models"
)

func setupTest() {
	// Database
	if models.DB == nil {
		models.InitializeDB("test")
	}

	models.DB.Exec("TRUNCATE reviews")
	models.DB.Exec("TRUNCATE scores")

	// Clients
	if PlacesClient == nil {
		PlacesClient = &clients.MockPlacesClient{}
	}

}
