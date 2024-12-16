package main

import (
	"log"
	"net/http"

	"processout-coding-challenge-tash-canter/handlers"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/payments", handlers.ProcessPaymentHandler).Methods("POST")
	r.HandleFunc("/payments/{id}", handlers.GetPaymentDetailsHandler).Methods("GET")
	return r
}

func main() {
    r := setupRouter()

	log.Println("Server starting on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
