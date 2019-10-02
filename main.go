package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	"github.com/kiambogo/coffeeworks/clients"
)

var Env Environment
var PlacesClient clients.PlacesIface

func main() {
	port := 8080

	loadEnv()
	initPlacesClient()
	r := BuildRoutes()

	log.Printf("CoffeeWorks started in environment '%v'", Env)
	log.Printf("Listening on %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), handlers.CORS()(r))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")
	switch env {
	case "test":
		Env = Test
	case "development":
		Env = Development
	case "production":
		Env = Production
	default:
		Env = Test
	}
}

func initPlacesClient() {
	key := os.Getenv("PLACES_API_KEY")
	if key == "" && Env == Test {
		log.Fatal("PLACES_API_KEY is not specified")
	}

	if Env == Development || Env == Production {
		PlacesClient = clients.InitializePlacesClient(key)
	} else {
		PlacesClient = &clients.MockPlacesClient{}
	}
}

type Environment string

const Production = Environment("production")
const Development = Environment("development")
const Test = Environment("test")
