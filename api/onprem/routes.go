package main

import (
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/api/onprem/handlers"
	"github.com/gorilla/mux"
)

func routes(handlers handlers.Handler) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/stock", handlers.AddStockHandler).Methods("POST")
	root.HandleFunc("/stock", handlers.ListStockHandler).Methods("GET")
	root.HandleFunc("/stock/{stockId}", handlers.DeleteStockHandler).Methods("DELETE")

	return root
}
