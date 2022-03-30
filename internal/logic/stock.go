package logic

import (
	"context"

	"github.com/MauricioAntonioMartinez/mcbot/internal/shared"
)

func (l Logic) AddStock(ctx context.Context, params shared.AddStockParams) error {
	return nil
}

func (l Logic) AnalyseStock(ctx context.Context, params shared.AnalyseStockParams) (shared.AnalyseStockResult, error) {
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
	return shared.AnalyseStockResult{}, nil
}
