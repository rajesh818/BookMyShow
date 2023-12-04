package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/models"
	"gorm.io/gorm"
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

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	log.Println("Get user by email started")
	vars := mux.Vars(r)
	email := vars["email"]
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	email := queryParams.Get("email")
	id := queryParams.Get("id")
	var result *gorm.DB
	var user models.User
	if id == "" && email == "" {
		http.Error(w, "No email or Id available to fetch data", http.StatusBadRequest)
		return
	}
	if id == "" {
		result = DB.Where("email = ?", email).First(&user)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		result = DB.Where("id = ?", id).First(&user)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
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

func DeleteUser(w http.ResponseWriter, r *http.Request){
	log.Println("GetUserById Function reached")
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	result := DB.Where("id = ?", id).Delete(&user);
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
