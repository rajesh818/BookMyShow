package routes

import (
	"fmt"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	fmt.Println("Registering Routes")
	BasicRoutes(r);
	UserRoutes(r)
}
