package handlers

import (
	"context"
	"fmt"

	"github.com/mcuv3/mcbot/internal/collectors/binance"
)

type KLineHandler struct{}

func (k KLineHandler) OnMessage(ctx context.Context, msg binance.KlinePayload) error {
	fmt.Println("Message : ", msg)
	return nil
}

func (k KLineHandler) OnError(ctx context.Context, err error) {
	fmt.Println("error : ", err)
}

func (k KLineHandler) OnStart(msg []byte) {
	fmt.Println("Start : ", string(msg))
}
