package handlers

import (
	"net/http"

	"github.com/mcuv3/mcbot/internal/shared"
)

type analyzeStockRequest struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
}

func (l Handlers) AnalyzeStockHandler(w http.ResponseWriter, r *http.Request) {
	params := shared.AnalyzeStockParams{}
	l.parseRequest(r, &params)
	res, err := l.logic.AnalyzeStock(r.Context(), params)
	if err != nil {
		l.writeError(w, err)
		return
	}
	l.writeSuccessResponse(w, res)
}
