package ingestors

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mcuv3/mcbot/internal/storage"
	"github.com/mcuv3/mcbot/internal/storage/kline"
)

type Gartley struct {
	storage.Store
	process []ingest
}

func NewGartley(stores storage.Store) *Gartley {
	return &Gartley{
		process: []ingest{},
		Store:   stores,
	}
}

type Provider interface {
	nextKline(ctx context.Context, s string) (kline.Model, error)
}

type ingest struct {
	Provider
	symbol    string
	process   chan string
	buffKline []kline.Model
}

func (i *ingest) Start(ctx context.Context) {
	for {
		fmt.Printf("Waiting for kline ...  symbol:%s \n", i.symbol)
		kline, err := i.nextKline(ctx, i.symbol)
		if err != nil {
			continue
		}
		if len(i.buffKline) == 0 {
			i.buffKline = append(i.buffKline, kline)
			continue
		}

		if i.buffKline[len(i.buffKline)-1].CreatedAt >= kline.CreatedAt {
			continue
		}

		i.buffKline = append(i.buffKline, kline)

		if len(i.buffKline) < 30 {
			continue // not required the minimun amount of kline to start ingesting.
		}
		points, err := findMaxAndMinPoints(i.buffKline, 25)
		if err != nil {
			log.Println("Error finding max and min points:", err)
			continue
		}
		fmt.Println("Points: ", points)

		time.Sleep(time.Minute) // wait unitil the new kline gets stored.
	}
	// calculate fibo and start
}

func findMaxAndMinPoints(elements []kline.Model, perChange float64) ([]float64, error) {
	if len(elements) < 2 {
		return nil, errors.New("not enough elements to find max and min")
	}
	if perChange < 0 || perChange > 100 {
		return nil, errors.New("perChange must be between 0 and 100")
	}
	points := []float64{}
	points = append(points, elements[0].HighPrice)
	up := elements[0].HighPrice > elements[1].HighPrice

	add := func(v float64, replace bool) {
		if replace {
			points[len(points)-1] = v
		} else {
			last := points[len(points)-1]
			per := math.Abs(100 - ((v * 100) / last))
			if per >= perChange {
				// should change
				up = !up
				points = append(points, v)
			}
		}
	}

	for _, e := range elements {
		if e.HighPrice > points[len(points)-1] {
			add(e.HighPrice, up)
		} else {
			add(e.HighPrice, !up)
		}
	}
	return points, nil

}

func (g *Gartley) IngestSymbol(ctx context.Context, s string) {
	ingest := ingest{
		Provider:  g,
		symbol:    s,
		process:   make(chan string),
		buffKline: []kline.Model{},
	}
	go ingest.Start(ctx)
	g.process = append(g.process, ingest)
}

func (g *Gartley) nextKline(ctx context.Context, s string) (kline.Model, error) {
	k, err := g.Kline.GetLast(ctx, s)
	if err != nil {
		return kline.Model{}, err
	}
	return k, nil
}
