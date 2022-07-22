package ingestors

import (
	"context"
	"math"
	"time"

	"github.com/mcuv3/mcbot/internal/storage"
	"github.com/mcuv3/mcbot/internal/storage/kline"
)

type Gartley struct {
	storage.Store
	process []ingest
}

type Provider interface {
	NextKline(ctx context.Context, s string) (kline.Model, error)
}

type ingest struct {
	Provider
	symbol    string
	process   chan string
	buffKline []kline.Model
}

func (i *ingest) Start(ctx context.Context) {
	go func() {

		for {
			kline, err := i.NextKline(ctx, i.symbol)
			if err != nil {
				continue
			}
			i.buffKline = append(i.buffKline, kline)

			if len(i.buffKline) < 30 {
				continue // not required the minimun amount of kline to start ingesting.
			}

			time.Sleep(time.Minute) // wait unitil the new kline gets stored.
		}
		// calculate fibo and start
	}()
}

func findInfexionPoints(elments []kline.Model) []float64 {
	points := []float64{}
	points = append(points, elments[0].HighPrice)
	up := elments[0].HighPrice > elments[1].HighPrice

	preferedPer := 2

	add := func(v float64, replace bool) {
		if replace {
			points[len(points)-1] = v
		} else {
			last := points[len(points)-1]
			diff := math.Abs(last - v)
			if diff >= float64(preferedPer) {
				// should change
				up = !up
				points = append(points, v)
			}
		}
	}

	for _, e := range elments {
		if e.HighPrice > points[len(points)-1] {
			add(e.HighPrice, up)
		} else {
			add(e.HighPrice, !up)
		}
	}
	return points

}

func (g *Gartley) IngestSymbol(ctx context.Context, s string) {
	ingest := ingest{
		Provider:  g,
		symbol:    s,
		process:   make(chan string),
		buffKline: []kline.Model{},
	}
	ingest.Start(ctx)
	g.process = append(g.process, ingest)
}

func (g *Gartley) NextKline(ctx context.Context, s string) (kline.Model, error) {
	return kline.Model{}, nil
}
