package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/models"
)

type LocationDetails struct{
	City string
	State string
	Pincode int 
}

func CreateLocation(w http.ResponseWriter, r *http.Request){
	var requestBody LocationDetails
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "No Request body", http.StatusInternalServerError)
		return
	}
	location := models.Location{
		City: requestBody.City,
		State: requestBody.State,
		Pincode: requestBody.Pincode,
	}
	result := DB.Create(&location);
	if(result.Error != nil){
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(location); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func GetLocationById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["locationId"]
	var location models.Location
	result := DB.Where("id = ?", id).First(&location);
	if(result.Error != nil){
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(location); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func DeleteLocationById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["locationId"]
	var location models.Location
	result := DB.Where("id = ?", id).Delete(&location);
	if(result.Error != nil){
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(location); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func UpdateLocationById(w http.ResponseWriter, r *http.Request){
	var requestBody LocationDetails
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "No Request body", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id := vars["locationId"]
	var location models.Location
	result := DB.Where("id = ?", id).First(&location);
	if(result.Error != nil){
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	location.City = requestBody.City;
	location.State = requestBody.State;
	location.Pincode = requestBody.Pincode;
	DB.Save(&location);
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(location); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

