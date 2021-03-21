package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihirgupta0900/go-petshop/database"
	"github.com/mihirgupta0900/go-petshop/models"
)

func AddUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", UsersHandler).Methods("GET")
	r.HandleFunc("/users", AddUsershandler).Methods("POST")
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var pets []models.Pet
	database.DB.Find(&pets)

	fmt.Println("hey")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func AddUsershandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
