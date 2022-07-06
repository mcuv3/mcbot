package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcuv3/mcbot/internal/shared"
	"github.com/mcuv3/mcbot/internal/storage/stock"
)

func (l Handlers) AddStockHandler(w http.ResponseWriter, r *http.Request) {
	req := analyzeStockRequest{}
	if err := l.parseRequest(r, &req); err != nil {
		l.writeError(w, err)
		return
	}

	err := l.logic.AddStock(r.Context(), shared.AddStockParams{
		Symbol:   req.Symbol,
		Exchange: req.Exchange,
	})

	if err != nil {
		l.writeError(w, err)
		return
	}

	l.writeSuccessResponse(w, map[string]string{
		"message": fmt.Sprintf("Stock %s added successfully", req.Symbol),
	})
}

func (l Handlers) ListStockHandler(w http.ResponseWriter, r *http.Request) {
	stocks, err := l.logic.ListStocks(context.Background(), stock.ListParams{})
	if err != nil {
		l.writeError(w, err)
		return
	}

	l.writeSuccessResponse(w, stocks)
}

func (l Handlers) DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, ok := params["stockId"]
	if !ok {
		l.writeError(w, fmt.Errorf("stockId is required"))
		return
	}

	err := l.logic.DeleteStock(r.Context(), stockId)
	if err != nil {
		l.writeError(w, err)
		return
	}

	l.writeSuccessResponse(w, map[string]string{
		"message": fmt.Sprintf("Stock ID:%s deleted successfully", stockId),
	})
}
