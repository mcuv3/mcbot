package shared

type Analyser interface {
	Analyse(AnalyserParams) ([]Trend, error)
}

type AnalyserParams struct {
	Frame     int // number of days to analyse and get trends of it
	TrendSize     // Amount of hours to look for a trend
	Symbol    string
	Collector Collector
}
