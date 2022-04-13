package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/cmd/handleFunc"
)

const (
	GET = "GET"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"

	//Health check server
	router.HandleFunc("/health", handleFunc.Health).Methods(GET)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("server is not ready!")
	}
}
