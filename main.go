package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/routes"
	"github.com/rajesh818/BookMyShow/api/handlers"
	"github.com/rajesh818/BookMyShow/database"
)

func main() {
	fmt.Println("Server Started")

	db, err := database.Connect()
	if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }
	handlers.SetDatabase(db);
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8083",router))
}
