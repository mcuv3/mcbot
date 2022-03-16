// run it on your machine as a web server instead of lambdas, create orders,feed data with grpc calls
package main

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/mcbot/internal/analysis"
	"github.com/MauricioAntonioMartinez/mcbot/internal/collectors"
	"github.com/MauricioAntonioMartinez/mcbot/internal/shared"
)

func main() {

	cryptoCollector := collectors.NewCryptoStore()

	an := analysis.NewAnalyser(shared.AnalyserParams{
		Frame:     365 * 2, // four years
		Collector: cryptoCollector,
		Symbol:    "BTC",
		TrendSize: 24,
	})
	trends, err := an.Analyse()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(trends))

	for _, v := range trends {
		fmt.Printf("%+v\n", v)
	}

}
