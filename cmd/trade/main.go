package main

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/mcbot/analysis"
	"github.com/MauricioAntonioMartinez/mcbot/collectors"
	"github.com/MauricioAntonioMartinez/mcbot/shared"
)

func main() {

	cryptoCollector := collectors.NewCryptoStore()

	an := analysis.NewAnalyser(shared.AnalyserParams{
		Frame:     365 * 2, // four years
		Collector: cryptoCollector,
		Symbol:    "BTC",
		TrendSize: 2,
	})
	_, err := an.Analyse()
	if err != nil {
		fmt.Println(err)
		return
	}

}
