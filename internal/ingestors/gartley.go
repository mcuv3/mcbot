package ingestors

import (
	"context"
	"time"

	"github.com/mcuv3/mcbot/internal/storage"
	"github.com/mcuv3/mcbot/internal/storage/kline"
)

type Gartley struct {
	storage.Store
	process []ingest
}

type Provider interface {
	GetKline(s string) kline.Model
}

type ingest struct {
	Provider
	symbol    string
	process   chan string
	buffKline []kline.Model
}

func (i *ingest) Start() {
	go func() {

		for {
			kline := i.GetKline(i.symbol)
			i.buffKline = append(i.buffKline, kline)

			if len(i.buffKline) < 30 {
				continue // not required the minimun amount of kline to start ingesting.
			}

			time.Sleep(time.Minute) // wait unitil the new kline gets stored.
		}
		// calculate fibo and start
	}()
}

func (g *Gartley) IngestSymbol(ctx context.Context, s string) {
	ingest := ingest{
		symbol:    s,
		process:   make(chan string),
		buffKline: []kline.Model{},
	}
	ingest.Start()
	g.process = append(g.process, ingest)
}

func (g *Gartley) GetKline(ctx context.Context, s string) (kline.Model, error) {
	return kline.Model{}, nil
}
