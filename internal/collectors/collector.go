package collectors

import (
	"github.com/MauricioAntonioMartinez/mcbot/internal/collectors/coinapi"
	"github.com/MauricioAntonioMartinez/mcbot/internal/collectors/cryptocomp"
	"github.com/MauricioAntonioMartinez/mcbot/internal/shared"
)

type CryptoStore struct {
	collectors map[string]shared.Collector
}

func NewCryptoStore() shared.Collector {
	return CryptoStore{
		collectors: map[string]shared.Collector{
			"cryptocomp": cryptocomp.NewCryptoComp(""),
			"coinapi":    coinapi.NewCoinAPI(),
		},
	}
}

func (c CryptoStore) Collect(symbol string, frame int) ([]shared.Candle, error) {
	for _, v := range c.collectors {
		candles, err := v.Collect(symbol, frame)
		if err != nil {
			continue
		}
		return candles, nil
	}
	return nil, nil
}
