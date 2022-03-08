package main

import (
	"fmt"

	"gopkg.in/resty.v0"
)

func main() {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", "xxx").
		Get("https://rest.coinapi.io/v1/quotes/BITSTAMP_SPOT_BTC_USD/history?time_start=2016-01-01T00:00:00&time_end=2017-01-01T00:00:00")
	fmt.Println(resp, err)
}
