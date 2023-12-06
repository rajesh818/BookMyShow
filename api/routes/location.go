package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/handlers"
)

func LocationRoutes(r *mux.Router) {
	log.Println("LocationRoutes")
	// Create a location - POST
	r.HandleFunc("/create", handlers.CreateLocation).Methods(http.MethodPost)

	// Get a location - GET
	r.HandleFunc("/get/{locationId}", handlers.GetLocationById).Methods(http.MethodGet);

	// Delete a location - DELETE
	r.HandleFunc("/delete/{locationId}", handlers.DeleteLocationById).Methods(http.MethodDelete);

	// Update a location - PUT
	r.HandleFunc("/update/{locationId}", handlers.UpdateLocationById).Methods(http.MethodPut);
}