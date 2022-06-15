package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcuv3/mcbot/api/onprem/handlers"
)

func routes(handlers handlers.Handler) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/stock", handlers.AddStockHandler).Methods("POST")
	root.HandleFunc("/stock", handlers.ListStockHandler).Methods("GET")
	root.HandleFunc("/stock/{stockId}", handlers.DeleteStockHandler).Methods("DELETE")
	root.HandleFunc("/trades", handlers.AddStatusHandler).Methods("GET")
	root.HandleFunc("/trades/current", handlers.CurrentStatusHandler).Methods("GET")
	root.HandleFunc("/seed", handlers.SeedHandler).Methods("GET")

	return root
}
