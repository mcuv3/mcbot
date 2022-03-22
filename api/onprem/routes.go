package main

import (
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/api/onprem/handlers"
	"github.com/gorilla/mux"
)

func routes(handlers handlers.Handler) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/analyse", handlers.AnalyseStockHandler).Methods("GET")

	return root
}
