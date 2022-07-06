package logic

import (
	"context"
	"strings"

	"github.com/mcuv3/mcbot/internal/collectors/binance"
	"github.com/mcuv3/mcbot/internal/storage/stock"
)

func (l Logic) Seed(ctx context.Context) error {
	client := binance.NewClient()
	symbols, err := client.ListSymbols(ctx)
	if err != nil {
		return err
	}

	stocks := make([]interface{}, 0, len(symbols))

	for _, s := range symbols {
		stocks = append(stocks, stock.Stock{
			Symbol:     strings.ToLower(s.Symbol),
			BaseAsset:  s.BaseAsset,
			QuoteAsset: s.QuoteAsset,
			Exchange:   "binance",
		})
	}

	return l.Stock.SaveMany(stocks)
}
