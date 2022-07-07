package logic

import (
	"context"
	"math/rand"

	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mcuv3/mcbot/internal/collectors/binance"
	"github.com/mcuv3/mcbot/internal/collectors/binance/handlers"
	"github.com/mcuv3/mcbot/internal/shared"
	"github.com/mcuv3/mcbot/internal/storage/stock"
)

func (l Logic) AddStock(ctx context.Context, params shared.AddStockParams) error {
	err := l.Stock.Save(stock.Stock{
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
	err := l.Stock.Delete(ctx, stock.DeleteParams{
		StockID: uuid.MustParse(stockID),
	})
	if err != nil {
		return err
	}
	return nil
}

func (l Logic) ListStocks(ctx context.Context, params stock.ListParams) ([]stock.Stock, error) {
	return l.Stock.Find(ctx, params)
}

func (l Logic) AnalyzeStock(ctx context.Context, params shared.AnalyzeStockParams) (shared.AnalyzeStockResult, error) {
	sb := binance.NewSymbolSubscriber[binance.KlinePayload]("", handlers.KLineHandler{
		Store: l.Store,
	})
	_, err := l.Stock.FindOne(ctx, stock.FindOneParams{
		Symbol:   params.Symbol,
		Exchange: params.Exchange,
	})
	if err != nil {
		return shared.AnalyzeStockResult{}, errors.New("couldn't find symbol for the given exchange")
	}

	go func() {
		if err := sb.Dial(ctx, binance.Params{
			ID:     rand.Intn(100),
			Method: binance.SUBSCRIBE,
			Params: []string{fmt.Sprintf("%s@kline_1m", params.Symbol)},
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
