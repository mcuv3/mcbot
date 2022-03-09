package shared

type Analyser interface {
	Analyse(AnalyserParams) ([]Trend, error)
}

type AnalyserParams struct {
	Frame int // number of days to analyse and get trends of it
	TrendSize
	Symbol    string
	Collector Collector
}
