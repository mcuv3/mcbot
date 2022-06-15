package logic

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mcuv3/mcbot/internal/collectors/binance"
	"github.com/mcuv3/mcbot/internal/collectors/binance/handlers"
	"github.com/mcuv3/mcbot/internal/shared"
	"github.com/mcuv3/mcbot/internal/storage/stock"
)

func (l Logic) AddStock(ctx context.Context, params shared.AddStockParams) error {
	err := l.SaveStock(stock.Stock{
		Symbol:    params.Symbol,
		BaseAsset: params.Symbol,
		Exchange:  params.Exchange,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l Logic) DeleteStock(ctx context.Context, stockID string) error {
	err := l.stores.Stock.Delete(ctx, stock.DeleteParams{
		StockID: uuid.MustParse(stockID),
	})
	if err != nil {
		return err
	}
	return nil
}

func (l Logic) ListStocks(ctx context.Context, params stock.ListParams) ([]stock.Stock, error) {
	return l.stores.Stock.Find(ctx, params)
}

func (l Logic) AnalyzeStock(ctx context.Context, params shared.AnalyzeStockParams) (shared.AnalyzeStockResult, error) {
	sb := binance.NewSymbolSubscriber[binance.KlinePayload]("", handlers.KLineHandler{})
	go func() {
		if err := sb.Dial(ctx, binance.Params{
			ID:     1,
			Method: binance.SUBSCRIBE,
			Params: []string{""},
		}); err != nil {
			log.Println(err)
		}
	}()

	return shared.AnalyzeStockResult{
		Symbol: params.Symbol,
	}, nil
}

//cryptoCollector := collectors.NewCryptoStore()

// an := analysis.NewAnalyser(shared.AnalyserParams{
// 	Frame:     365 * 2, // four years
// 	Collector: cryptoCollector,
// 	Symbol:    "BTC",
// 	TrendSize: 24,
// })
// trends, err := an.Analyse()
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(len(trends))

// for _, v := range trends {
// 	fmt.Printf("%+v\n", v)
// }
