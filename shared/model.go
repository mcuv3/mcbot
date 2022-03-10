package shared

type Exchange = string

const (
	Binance Exchange = "Binance"
)

type TrendState = int

const (
	IDLE TrendState = iota
	UP
	DOWN
)

type TrendSize int

const (
	Hour TrendSize = iota
	Day
	Week
	Month
)

type Stock struct {
	Symbol string
	Name   string
	Exchange
}

// Price is the current price of the stock at stream RPC call
type Price struct {
	Bid float64
	Ask float64
}

type Candle struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
	Time   string
}

type Trend struct {
	Avg        float64    // represents the average of the last n prices of the trend size
	State      TrendState // compared to the previous trend.
	Volatility float64    // factor of change
	Size       TrendSize
	Volume     float64 // not available at hourly trend.
}
