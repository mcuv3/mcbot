package handlers

import (
	"fmt"
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/internal/storage/stock"
	"github.com/google/uuid"
)

func (l Logic) AnalyseStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))

	err := l.Stores.Stock.Save(stock.Stock{
		Symbol:   "AAPL",
		ID:       uuid.New().String(),
		Name:     "Apple",
		Exchange: "NASDAQ",
	})

	if err != nil {
		fmt.Println(err)
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

}
