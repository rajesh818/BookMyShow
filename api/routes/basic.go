package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/handlers"
)

func BasicRoutes(r *mux.Router) {
	fmt.Println("Reached Router");
	r.HandleFunc("/", handlers.BasicHandler).Methods(http.MethodGet)
}
