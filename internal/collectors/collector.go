package collectors

import (
	"github.com/mcuv3/mcbot/internal/collectors/cryptocomp"
	"github.com/mcuv3/mcbot/internal/shared"
)

type CryptoStore struct {
	collectors map[string]shared.Collector
}

func NewCryptoStore() shared.Collector {
	return CryptoStore{
		collectors: map[string]shared.Collector{
			"cryptocomp": cryptocomp.NewCryptoComp(""),
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
