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
	Symbol   string
	Interval string
	Limit    int
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

func (b API) GetKline(ctx context.Context, params GetKlineParams) ([]Kline, error) {
	var klines []Kline
	var results [][]interface{}
	if err := b.Fetch(ctx, FetchParams{}, &results); err != nil {
		return nil, err
	}

	for _, r := range results {
		k := Kline{
			Symbol: r[0].(string),
		}
		klines = append(klines, k)
	}

	return klines, nil
}

func (b API) ListSymbols(ctx context.Context) ([]string, error) {
	type response struct {
		Symbols []struct {
			Symbol string `json:"symbol"`
		} `json:"symbols"`
	}
	var resp response
	if err := b.Fetch(ctx, FetchParams{
		Method: "GET",
		URL:    b.URL("exchangeInfo"),
	}, &resp); err != nil {
		return nil, err
	}
	symbols := []string{}
	for _, s := range resp.Symbols {
		symbols = append(symbols, s.Symbol)
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
