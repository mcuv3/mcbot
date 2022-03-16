package shared

import "context"

type API interface {
	AddStock(context.Context, AddStockParams) error
	AnalyseStock(context.Context, AnalyseStockParams) (AnalyseStockResult, error)
	Buy(context.Context, BuyParams) error
	Sell(context.Context, SellParams) error
}

type AddStockParams struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
}

type BuyParams struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type SellParams struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type AnalyseStockParams struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Frame    int    `json:"frame"`
}

type AnalyseStockResult struct {
	Symbol string `json:"symbol"`
}
