package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/internal/storage/stock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (l Logic) AddStockHandler(w http.ResponseWriter, r *http.Request) {
	req := analyseStockRequest{}
	if err := l.parseRequest(r, &req); err != nil {
		l.writeError(w, err)
		return
	}

	err := l.SaveStock(stock.Stock{
		Symbol:   req.Symbol,
		ID:       uuid.New().String(),
		Name:     req.Name,
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

func (l Logic) ListStockHandler(w http.ResponseWriter, r *http.Request) {
	stocks, err := l.ListStocks(context.Background(), stock.ListParams{})
	if err != nil {
		l.writeError(w, err)
		return
	}

	l.writeSuccessResponse(w, stocks)
}

func (l Logic) DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, ok := params["stockId"]
	if !ok {
		l.writeError(w, fmt.Errorf("stockId is required"))
		return
	}

	err := l.DeleteStock(context.Background(), stock.DeleteParams{
		StockID: uuid.MustParse(stockId),
	})
	if err != nil {
		l.writeError(w, err)
		return
	}

	l.writeSuccessResponse(w, map[string]string{
		"message": fmt.Sprintf("Stock ID:%s deleted successfully", stockId),
	})
}
