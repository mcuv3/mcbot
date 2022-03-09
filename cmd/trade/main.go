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
		Frame:     365 * 4, // four years
		Collector: cryptoCollector,
	})
	fmt.Println(an.Analyse())
}
