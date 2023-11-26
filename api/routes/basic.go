package basic

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/handlers"
)

func BasicRoutes(r *mux.Router) {
	fmt.Println("Reached Router");
	r.HandleFunc("/", basic.BasicHandler).Methods(http.MethodGet)
}
