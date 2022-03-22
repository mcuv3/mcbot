package coinapi

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/mcbot/internal/shared"
	"gopkg.in/resty.v0"
)

type CoinAPI struct {
}

func NewCoinAPI() shared.Collector {
	return CoinAPI{}
}

func (c CoinAPI) Collect(symbols string, frame int) ([]shared.Candle, error) {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", "xxx").
		Get("https://rest.coinapi.io/v1/quotes/BITSTAMP_SPOT_BTC_USD/history?time_start=2016-01-01T00:00:00&time_end=2017-01-01T00:00:00")
	fmt.Println(resp, err)
	return nil, nil
}
