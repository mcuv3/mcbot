package cryptocomp

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/MauricioAntonioMartinez/mcbot/shared"
	"gopkg.in/resty.v0"
)

type CryptoComp struct {
	apikey string
}

func NewCryptoComp(apikey string) shared.Collector {
	return CryptoComp{apikey}
}

// response from cryptocomp might change in the future.
type response struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	} `json:"RateLimit"`
	Data struct {
		Aggregated bool     `json:"Aggregated"`
		TimeFrom   int      `json:"TimeFrom"`
		TimeTo     int      `json:"TimeTo"`
		Data       []record `json:"Data"`
	} `json:"Data"`
}

type record struct {
	Time             int     `json:"time"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Open             float64 `json:"open"`
	Volumefrom       float64 `json:"volumefrom"`
	Volumeto         float64 `json:"volumeto"`
	Close            float64 `json:"close"`
	ConversionType   string  `json:"conversionType"`
	ConversionSymbol string  `json:"conversionSymbol"`
}

// 24 *

func (c CryptoComp) getData(symbol string, ts string, limit int) (response, error) {
	uri := newUri()
	if ts != "" {
		uri.ts(ts)
	}
	url := uri.apikey(c.apikey).
		symbol(symbol).limit(limit).tsym("USD").build()

	resp, err := resty.R().
		Get(url)
	if err != nil {
		return response{}, err
	}

	r := response{}
	if err := json.Unmarshal(resp.Body, &r); err != nil {
		return r, err
	}

	return r, nil
}

func (c CryptoComp) Collect(symbol string, frame int) ([]shared.Candle, error) {
	hours := float64(frame * 24)
	candles := []shared.Candle{}
	calls := int(math.Ceil(hours / c.recordLimit()))
	ts := ""

	for i := 0; i < calls; i++ {
		r, err := c.getData(symbol, ts, 2000)
		if err != nil {
			return nil, err
		}

		records := r.Data.Data

		if len(records) > 0 {
			fmt.Println(records[0].Time)
			ts = fmt.Sprintf("%d", records[0].Time)
		}

		for _, r := range r.Data.Data {
			candles = append(candles, c.toCandle(r))
		}
	}

	return candles, nil
}

func (c CryptoComp) toCandle(r record) shared.Candle {
	return shared.Candle{
		Open:   r.Open,
		High:   r.High,
		Low:    r.Low,
		Close:  r.Close,
		Volume: r.Volumeto,
		Time:   strconv.Itoa(r.Time),
	}
}

// recordLimit represents the maximum number of records that can be returned, this
// could change with the time placed on a function for readability
func (c CryptoComp) recordLimit() float64 {
	return 2000
}
