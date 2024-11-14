package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// func Router() *mux.Router {}
func Router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", DeleteStock).Methods("DELETE", "OPTIONS")

	return secureHeaders(router)
}
