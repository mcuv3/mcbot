package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcuv3/mcbot/handlers"
)

func routes(handlers handlers.Handler) http.Handler {
	root := mux.NewRouter()
	//
	root.HandleFunc("/stock", handlers.AddStockHandler).Methods("POST")
	root.HandleFunc("/stock", handlers.ListStockHandler).Methods("GET")
	root.HandleFunc("/stock/{stockId}", handlers.DeleteStockHandler).Methods("DELETE")
	root.HandleFunc("/analyze/state/{analyzerId}", handlers.AnalyzeStateHandler).Methods("GET")
	root.HandleFunc("/analyze", handlers.AnalyzeStockHandler).Methods("POST")
	root.HandleFunc("/seed", handlers.SeedHandler).Methods("GET")

	return root
}
