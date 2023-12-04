package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/models"
	"log"
	"net/http"
)

type UserDetails struct {
	Name     string
	Email    string
	Mobile   string
	Password string
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUserById Function reached")
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestBody UserDetails
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "No Request body", http.StatusInternalServerError)
		return
	}
	user := models.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Mobile:   requestBody.Mobile,
		Password: requestBody.Password,
	}

	result := DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}
