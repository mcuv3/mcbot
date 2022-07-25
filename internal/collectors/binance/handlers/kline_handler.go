package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/mcuv3/mcbot/internal/collectors/binance"
	"github.com/mcuv3/mcbot/internal/storage"
	"github.com/mcuv3/mcbot/internal/storage/kline"
)

type KLineHandler struct {
	storage.Store
}

func (k KLineHandler) OnMessage(ctx context.Context, msg binance.KlinePayload) error {
	fmt.Println("Message : ", msg)
	if msg.K.IsClosed {
		return k.Kline.Save(kline.Model{
			Symbol:     strings.ToLower(msg.Symbol),
			StartTime:  msg.K.StartTime,
			CloseTime:  msg.K.CloseTime,
			OpenPrice:  msg.K.OpenPrice,
			ClosePrice: msg.K.ClosePrice,
			HighPrice:  msg.K.HighPrice,
			LowPrice:   msg.K.LowPrice,
		})
	}
	return nil
}

func (k KLineHandler) OnError(ctx context.Context, err error) {
	fmt.Println("error : ", err)
}

func (k KLineHandler) OnStart(msg []byte) {
	fmt.Println("Start : ", string(msg))
}
