package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/handlers"
)

func UserRoutes(r *mux.Router) {
	log.Println("UserRoutes")
	r.HandleFunc("/create", handlers.CreateUser).Methods(http.MethodPost)

	// delete a user from db
	r.HandleFunc("/delete/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	// get user by email
	r.HandleFunc("/getByEmail/{email}", handlers.GetUserByEmail).Methods(http.MethodGet)

	// get user by id
	r.HandleFunc("/getById/{id}", handlers.GetUserById).Methods(http.MethodGet)

	// Get user by email/id
	r.HandleFunc("/get", handlers.GetUser).Methods(http.MethodGet)



}