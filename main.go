package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mihirgupta0900/go-petshop/database"
	"github.com/mihirgupta0900/go-petshop/models"
	"github.com/mihirgupta0900/go-petshop/routes"
)

func main() {
	log.SetPrefix("main: ")

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Connecting to DB
	db := database.InitDB()

	// Auto migration
	migrationErr := db.AutoMigrate(&models.User{}, &models.Pet{})
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	r := mux.NewRouter()

	routes.AddPetRoutes(r)
	routes.AddUserRoutes(r)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
