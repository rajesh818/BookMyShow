package handlers

import (
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
)

type BasicData struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

var DB *gorm.DB;

func SetDatabase(database *gorm.DB){
	DB = database
}

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	data := BasicData{
		Name: "Rajesh",
		Job:  "Software Engineer",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}