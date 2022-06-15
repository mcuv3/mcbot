package logic

import (
	"github.com/mcuv3/mcbot/internal/storage/stock"
)

// SaveStockHandler is a function that saves a stock
func (l Logic) SaveStock(stock stock.Stock) error {
	return l.stores.Stock.Save(stock)
}
