package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MauricioAntonioMartinez/mcbot/analysis"
	"github.com/MauricioAntonioMartinez/mcbot/shared"
)

func main() {

	start, err := time.Parse("2006-01-02", "2022-02-20")
	if err != nil {
		log.Fatal(err)
	}
	an := analysis.NewAnalyser(shared.AnalyserParams{
		Start: start,
		End:   time.Now(),
	})
	fmt.Println(an.Analyse())
}
