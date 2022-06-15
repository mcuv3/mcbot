package handlers

import (
	"net/http"
)

type analyzeStockRequest struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
}

func (l Handlers) AnalyzeStockHandler(w http.ResponseWriter, r *http.Request) {

}
