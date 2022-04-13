package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/cmd/handleFunc"
)

const (
	GET  = "GET"
	POST = "POST"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"

	//Health check server
	router.HandleFunc("/health", handleFunc.Health).Methods(GET)

	//AddProduct adding product in store
	router.HandleFunc("/product", handleFunc.Product).Methods(POST)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("server is not ready!")
	}
}
