package handlers

import (
	"context"

	"github.com/MauricioAntonioMartinez/mcbot/internal/storage/stock"
)

// SaveStockHandler is a function that saves a stock
func (l Logic) SaveStock(stock stock.Stock) error {
	return l.Stores.Stock.Save(stock)
}

func (l Logic) DeleteStock(ctx context.Context, params stock.DeleteParams) error {
	return l.Stores.Stock.Delete(ctx, params)
}

func (l Logic) ListStocks(ctx context.Context, params stock.ListParams) ([]stock.Stock, error) {
	return l.Stores.Stock.Find(ctx, params)
}
