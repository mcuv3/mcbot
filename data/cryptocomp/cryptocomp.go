package cryptocomp

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v0"
)

type Response struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	} `json:"RateLimit"`
	Data struct {
		Aggregated bool `json:"Aggregated"`
		TimeFrom   int  `json:"TimeFrom"`
		TimeTo     int  `json:"TimeTo"`
		Data       []struct {
			Time             int     `json:"time"`
			High             float64 `json:"high"`
			Low              float64 `json:"low"`
			Open             float64 `json:"open"`
			Volumefrom       float64 `json:"volumefrom"`
			Volumeto         float64 `json:"volumeto"`
			Close            float64 `json:"close"`
			ConversionType   string  `json:"conversionType"`
			ConversionSymbol string  `json:"conversionSymbol"`
		} `json:"Data"`
	} `json:"Data"`
}

func GetData(symbol string) (Response, error) {
	resp, err := resty.R().
		Get("https://min-api.cryptocompare.com/data/v2/histohour?fsym=BTC&tsym=USD&api_key=xxx")
	fmt.Println(resp, err)
	r := Response{}
	if err := json.Unmarshal(resp.Body, &r); err != nil {
		return r, err
	}

	return r, nil

}
