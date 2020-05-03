package api

import (
	"log"
	"os"

	"github.com/johnnyaustor/go-crud-jwt/api/controllers"
	"github.com/johnnyaustor/go-crud-jwt/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		log.Println("we are getting the env values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	seed.Load(server.DB)

	server.Run(":8080")
}