package shared

import "time"

type Analyser interface {
	Analyse(AnalyserParams) ([]Trend, error)
}

type AnalyserParams struct {
	Start time.Time
	End   time.Time
	TrendSize
	Symbol string
}
