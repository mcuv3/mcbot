// takes every other piece of the puzzle and decides wheather to sell or buy.
package logic

import (
	"context"

	"github.com/mcuv3/mcbot/internal/ingestors"
	"github.com/mcuv3/mcbot/internal/shared"
	"github.com/mcuv3/mcbot/internal/storage"
	"github.com/mcuv3/mcbot/internal/storage/stock"
)

var _ = Layer(&Logic{})

type Layer interface {
	AddStock(ctx context.Context, params shared.AddStockParams) error
	DeleteStock(ctx context.Context, stockID string) error
	ListStocks(ctx context.Context, params stock.ListParams) ([]stock.Stock, error)
	AnalyzeStock(ctx context.Context, params shared.AnalyzeStockParams) (shared.AnalyzeStockResult, error)
	Seed(ctx context.Context) error
	Buy(ctx context.Context, params shared.BuyParams) error
	Sell(ctx context.Context, params shared.SellParams) error
}

type Params struct {
	Stores  storage.Store
	Gartley *ingestors.Gartley
}

type Logic struct {
	storage.Store
	Gartley *ingestors.Gartley
}

func New(params Params) Logic {
	return Logic{
		params.Stores,
		params.Gartley,
	}
}
