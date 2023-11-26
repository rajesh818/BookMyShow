package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rajesh818/BookMyShow/api/routes"
)

func main() {
	fmt.Println("Server Started");
	router := mux.NewRouter();
	basic.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8083",router))
}
