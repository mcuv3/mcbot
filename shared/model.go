package shared

import (
	"time"

	"github.com/google/uuid"
)

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

// Amount of horus of a given trend
type TrendSize = int

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
	ID           uuid.UUID // trendID
	Avg          float64   // represents the average of the last n prices of the trend size
	Volatility   float64   // highest - lowest based on the last n candles
	Size         TrendSize
	AvgVolume    float64
	StdDeviation float64 // standard deviation of the last n prices of the trend size
	Start        time.Time
	End          time.Time
}
