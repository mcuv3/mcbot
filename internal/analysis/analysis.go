package analysis

import (
	"errors"
	"math"

	"github.com/google/uuid"
	"github.com/mcuv3/mcbot/internal/shared"
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

func (a *Analyzer) Analyse() (trends []shared.Trend, err error) {
	data, err := a.cryptostore.Collect(a.symbol, a.frame)
	if err != nil {
		return nil, err
	}

	currentCandle := make([]shared.Candle, 0)
	var count int
	var t *shared.Trend
	var trendID uuid.UUID

	for _, c := range data {
		currentCandle = append(currentCandle, c)
		count++
		if count == a.trendSize {
			t, err = a.getTrend(currentCandle, trendID)
			if err != nil {
				return nil, err
			}
			if len(trends) > 0 {
				trends[len(trends)-1].NextTrendID = t.ID
			}
			trends = append(trends, *t)
			currentCandle, count = nil, 0
		}
	}

	return trends, err
}

func (a *Analyzer) getTrend(data []shared.Candle, prevTrendId uuid.UUID) (*shared.Trend, error) {
	if len(data) == 0 {
		return nil, errors.New("Not enough candles")
	}
	t := shared.Trend{
		ID:    uuid.New(),
		Size:  a.trendSize,
		Start: shared.UnixToDate(data[0].Time),
		End:   shared.UnixToDate(data[len(data)-1].Time),
	}

	var avg, avgVol float64
	high := data[0].High
	low := data[0].Low

	for _, c := range data {
		avg += c.Open
		avgVol += c.Volume
		if c.High > high {
			high = c.High
		}
		if c.Low < low {
			low = c.Low
		}
	}

	avg = avg / float64(len(data))
	var deviation float64
	for _, c := range data {
		deviation += math.Pow((c.Open - avg), 2)
	}

	t.StdDeviation = math.Sqrt(deviation / float64(len(data)))
	t.Volatility = high - low
	t.Avg = avg
	t.AvgVolume = avgVol / float64(len(data))

	return &t, nil
}
