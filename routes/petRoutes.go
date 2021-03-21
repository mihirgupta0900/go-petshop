package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihirgupta0900/go-petshop/database"
	"github.com/mihirgupta0900/go-petshop/models"
)

func AddPetRoutes(r *mux.Router) {
	r.HandleFunc("/pets", PetsHandler).Methods("GET")
	r.HandleFunc("/pets", AddPetshandler).Methods("POST")
}

func PetsHandler(w http.ResponseWriter, r *http.Request) {
	var pets []models.Pet
	database.DB.Find(&pets)

	fmt.Println("hey")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func AddPetshandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var pet models.Pet

	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&pet)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}
