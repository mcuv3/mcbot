package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type FetchParams struct {
	Method  string
	URL     string
	Query   map[string]string
	Results int
	Body    interface{}
}

type API struct {
	Endpoint string
	Version  string
	client   http.Client
}

func NewClient() API {
	return API{
		Endpoint: "api.binance.com",
		Version:  "v3",
		client:   http.Client{},
	}
}

func (b API) URL(path string) string {
	u := url.URL{
		Host:   b.Endpoint,
		Path:   fmt.Sprintf("/api/%s/%s", b.Version, path),
		Scheme: "https",
	}
	return u.String()
}

type GetKlineParams struct {
	Symbol    string
	Interval  string
	startTime int64
}

type Kline struct {
	Symbol           string
	OpenTime         int
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	CloseTime        int
	QuoteAssetVolume string
	NumberOfTrades   int
}

func (b API) GetKlines(ctx context.Context, params GetKlineParams) ([]Kline, error) {
	var klines []Kline
	var results [][]interface{}
	if err := b.Fetch(ctx, FetchParams{
		Method: http.MethodGet,
		URL:    b.URL("klines"),
		Query: map[string]string{
			"symbol":    params.Symbol,
			"interval":  params.Interval,
			"startTime": fmt.Sprintf("%d", params.startTime),
		},
	}, &results); err != nil {
		return nil, err
	}

	for _, r := range results {
		k := Kline{
			Symbol:           params.Symbol,
			OpenTime:         r[0].(int),
			Open:             r[1].(string),
			High:             r[2].(string),
			Low:              r[3].(string),
			Close:            r[4].(string),
			Volume:           r[5].(string),
			CloseTime:        r[6].(int),
			QuoteAssetVolume: r[7].(string),
			NumberOfTrades:   r[8].(int),
		}
		klines = append(klines, k)
	}

	return klines, nil
}

type Symbol struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

func (b API) ListSymbols(ctx context.Context) ([]Symbol, error) {
	type response struct {
		Symbols []struct {
			Symbol     string `json:"symbol"`
			BaseAsset  string `json:"baseAsset"`
			QuoteAsset string `json:"quoteAsset"`
		} `json:"symbols"`
	}
	var resp response
	if err := b.Fetch(ctx, FetchParams{
		Method: "GET",
		URL:    b.URL("exchangeInfo"),
	}, &resp); err != nil {
		return nil, err
	}
	symbols := []Symbol{}
	for _, s := range resp.Symbols {
		symbols = append(symbols, s)
	}
	return symbols, nil
}

// Fetch fetches the given URL and decodes the response into the given dest
func (f *API) Fetch(ctx context.Context, params FetchParams, dest interface{}) error {
	var postBody io.Reader = nil
	// use a post body if it is set
	if params.Body != nil {
		byteBody, err := json.Marshal(params.Body)
		if err != nil {
			return err
		}
		postBody = strings.NewReader(string(byteBody))
	}
	req, err := http.NewRequestWithContext(ctx, params.Method, params.URL, postBody)
	if err != nil {
		return err
	}

	if len(params.Query) > 0 {
		q := req.URL.Query()

		for k, v := range params.Query {
			q.Add(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusInternalServerError {
		msg, _ := ioutil.ReadAll(resp.Body)
		log.Printf("service error: %s  error: %s ", resp.Status, msg)
		return fmt.Errorf("service error: %s", resp.Status)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		d := map[string]interface{}{}
		if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
			return err
		}
		log.Printf("client error: %s %s returned code %d & body: %v", params.Method, params.URL, resp.StatusCode, d)
		return fmt.Errorf("client error: %s %s returned code %d", params.Method, params.URL, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		log.Printf("Api fetch Error Decode err %s", err)
		return err
	}

	return nil
}
