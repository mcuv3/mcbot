package shared

import (
	"context"
)

type API interface {
	AddStock(context.Context, AddStockParams) error
	AnalyzeStock(context.Context, AnalyzeStockParams) (AnalyzeStockResult, error)
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

type AnalyzeStockParams struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
}

type AnalyzeStockResult struct {
	Symbol string `json:"symbol"`
}

type GraphAnalysisParams struct {
	AnalyzerID string `json:"analyzer_id"`
	Symbol     string `json:"symbol"`
	Exchange   string `json:"exchange"`
}
