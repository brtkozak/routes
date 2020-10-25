package main

import (
	"log"
	"net/http"

	"github.com/brtkozak/routes/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/routes", api.GetRoutes).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
