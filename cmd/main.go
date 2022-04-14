package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/cmd/handleFunc"
	"store/pkg/application/repository"
)

const (
	GET  = "GET"
	POST = "POST"
)

func main() {
	log.Println("Server start connect to DB")
	if err := repository.Connect(); err != nil {
		log.Fatal("Server can't connect to db, error is ", err)
	}

	router := mux.NewRouter()
	port := ":8080"

	//Health check server
	router.HandleFunc("/health", handleFunc.Health).Methods(GET)

	//AddProduct adding product in store
	router.HandleFunc("/product", handleFunc.Product).Methods(POST)

	//GetAllProducts
	//router.HandleFunc("/allProducts", handleFunc.AllProducts).Methods(GET)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("server is not ready!")
	}
}
