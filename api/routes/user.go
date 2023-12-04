package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/handlers"
)

func UserRoutes(r *mux.Router) {
	log.Println("UserRoutes")
	r.HandleFunc("/get/{id}", handlers.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/create", handlers.CreateUser).Methods(http.MethodPost)
}