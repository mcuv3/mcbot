package binance

import (
	"context"
	"fmt"
)

type KLineHandler struct{}

func (k KLineHandler) OnMessage(ctx context.Context, msg KlinePayload) error {
	fmt.Println("Message : ", msg)
	return nil
}

func (k KLineHandler) OnError(ctx context.Context, err error) {
	fmt.Println("error : ", err)
}

func (k KLineHandler) OnStart(msg []byte) {
	fmt.Println("Start : ", string(msg))
}
