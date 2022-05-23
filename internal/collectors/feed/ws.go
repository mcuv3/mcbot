package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type WsHandler[T Payloads] interface {
	OnMessage(ctx context.Context, msg T) error
	OnError(ctx context.Context, err error)
}

const DEFAULT_ADDR = "stream.binance.com:9443"

type SymbolSubscriber[T Payloads] struct {
	addr    string
	handler WsHandler[T]
}

type Params struct {
	ID     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

func NewSymbolSubscriber[T Payloads](addr string, handler WsHandler[T]) SymbolSubscriber[T] {
	if addr == "" {
		addr = DEFAULT_ADDR
	}
	var sb SymbolSubscriber[T] = SymbolSubscriber[T]{
		addr:    addr,
		handler: handler,
	}
	return sb
}

func (w SymbolSubscriber[T]) Dial(ctx context.Context, params Params) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: w.addr, Path: "/ws"}

	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer c.Close()
	done := make(chan struct{})

	go func() {
		err := c.WriteMessage(websocket.TextMessage, payload)
		fmt.Println(err)
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				w.handler.OnError(ctx, err)
				return
			}
			hold := new(T)
			if err := json.Unmarshal(message, hold); err != nil {
				w.handler.OnError(ctx, err)
				return
			}
			w.handler.OnMessage(ctx, *hold)
		}
	}()

	for {
		select {
		case <-done:
			return nil
		case <-interrupt:
			log.Println("interrupt signal received. Exiting...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				w.handler.OnError(ctx, err)
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}
