package shared

type Collector interface {
	Collect(symbol string, frame int) ([]Candle, error)
}
