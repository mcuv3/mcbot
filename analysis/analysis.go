package analysis

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/mcbot/shared"
)

type Analyzer struct {
	frame       int
	symbol      string
	trendSize   shared.TrendSize
	cryptostore shared.Collector
}

func NewAnalyser(params shared.AnalyserParams) *Analyzer {
	return &Analyzer{
		frame:       params.Frame,
		symbol:      params.Symbol,
		trendSize:   params.TrendSize,
		cryptostore: params.Collector,
	}
}

func (a *Analyzer) Analyse() ([]shared.Trend, error) {
	data, err := a.cryptostore.Collect(a.symbol, a.frame)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(data))
	return nil, nil
}
