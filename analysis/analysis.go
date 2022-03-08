package analysis

import (
	"fmt"
	"time"

	"github.com/MauricioAntonioMartinez/mcbot/shared"
	"github.com/markcheno/go-quote"
)

type Analyzer struct {
	start  time.Time
	end    time.Time
	symbol string
	size   shared.TrendSize
}

func NewAnalyser(params shared.AnalyserParams) *Analyzer {
	return &Analyzer{
		start:  params.Start,
		end:    params.End,
		symbol: params.Symbol,
		size:   params.TrendSize,
	}
}

func (a *Analyzer) Analyse() ([]shared.Trend, error) {
	data, err := a.collect(a.start, a.end)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(data) / a.getSize())
	return nil, nil
}

func (a *Analyzer) collect(start, end time.Time) ([]shared.Candle, error) {
	_start, _end := start.Format("2006-01-02 15:04"), end.Format("2006-01-02 15:04")
	spy, err := quote.NewQuoteFromYahoo("spy", _start, _end, quote.Hour4, true)
	if err != nil {
		return nil, err
	}

	prices := []shared.Candle{}
	for idx, p := range spy.Close {
		prices = append(prices, shared.Candle{
			Close: p,
			Open:  spy.Open[idx],
			High:  spy.High[idx],
			Low:   spy.Low[idx],
		})
	}
	return prices, nil
}

func (a *Analyzer) getSize() int {
	switch a.size {
	case shared.Hour:
		return 4
	case shared.Day:
		return 6
	case shared.Week:
		return 6 * 7
	case shared.Month:
		return 6 * 30
	default:
		return 6
	}
}
