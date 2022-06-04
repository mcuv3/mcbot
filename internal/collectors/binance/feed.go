package binance

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Feed struct {
	Symbol string
}

func New() Feed {
	return Feed{}
}

func (f Feed) Start() {
	h := KLineHandler{}
	s := NewSymbolSubscriber[KlinePayload](DEFAULT_ADDR, h)

	if err := s.Dial(context.Background(), Params{
		ID:     1,
		Method: "SUBSCRIBE",
		Params: []string{"btcusdt@kline_1m", "ethusdt@kline_1m"},
	}); err != nil {
		log.Error().Err(err).Msg("failed to subscribe")
	}
}
