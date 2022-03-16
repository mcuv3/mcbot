package logic

import (
	"context"

	"github.com/MauricioAntonioMartinez/mcbot/internal/shared"
)

func (l Logic) AddStock(ctx context.Context, params shared.AddStockParams) error {
	return nil
}

func (l Logic) AnalyseStock(ctx context.Context, params shared.AnalyseStockParams) (shared.AnalyseStockResult, error) {
	return shared.AnalyseStockResult{}, nil
}
