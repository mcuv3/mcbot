package ingestors

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mcuv3/mcbot/internal/shared"
	"github.com/mcuv3/mcbot/internal/storage"
	"github.com/mcuv3/mcbot/internal/storage/kline"
)

var store storage.Store

type Gartley struct {
	store   storage.Store
	process []ingest
}

func NewGartley(stores storage.Store) *Gartley {
	store = stores
	fmt.Println(store.Kline.GetLast(context.Background(), "btcbusd"))

	return &Gartley{
		process: []ingest{},
		store:   stores,
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

type Armonic struct {
	X Point
	A Point
	B Point
	C Point
	D Point
}

func (i *ingest) Start(ctx context.Context) {
	for {
		fmt.Printf("Waiting for kline ...  symbol:%s \n", i.symbol)
		time.Sleep(time.Second * 5)
		kline, err := i.nextKline(ctx, i.symbol)
		if err != nil {
			log.Println(err)
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

		if len(i.buffKline) < 5 {
			continue // not required the minimun amount of kline to start ingesting.
		}

		max, min := FindMaxAndMinPoint(i.buffKline)
		pattern := Armonic{
			X: min,
			A: max,
		}

		log.Println("Points: ", max, min)

		rts := shared.GetFibonacciRetrace(min.Value, max.Value)

		var B float64
		var inTheBox bool
		var idxStartForC int

		for i, k := range i.buffKline[max.Index:] {
			if k.GetClosePrice() >= rts.L1618 && k.GetClosePrice() <= rts.L786 {
				if B == 0 || k.GetClosePrice() < B {
					if B != 0 {
						inTheBox = true
					}
					pattern.B = Point{
						Value: k.GetClosePrice(),
						Index: i + max.Index,
					}
				}
			} else if inTheBox { // we enter into the box but exited after
				idxStartForC = i
				break
			}
		}
		_ = idxStartForC

		// wait unitil the new kline gets stored.

		// a) Define max and min points knowledge X
		// b) Trace x - a fibonacci retrace Grace range [618 - 786] = b
		// c) Trace a - b fibonacci retrace Grace range [618 or above but not greater than a] = c
		// d) Trace b - a (opposite dir)
		// min 786 trace x-a if is below we take it and VALIDATE
		//RANGE x > 1 270  &&  x < 1 618
		//
	}
}

type Point struct {
	Value float64
	Index int
}

func FindMaxAndMinPoint(elements []kline.Model) (max Point, min Point) {
	for i, k := range elements {
		if k.GetHighPrice() > max.Value {

			max = Point{
				Value: k.GetOpenPrice(),
				Index: i,
			}
		}
		if min.Value == 0 || k.GetClosePrice() < min.Value {

			min = Point{
				Value: k.GetClosePrice(),
				Index: i,
			}
		}
	}
	return
}

// DEPRECATED: our current implementation doesn't need this :(
func FindMaxAndMinPoints(elements []kline.Model, perChange float64) ([]float64, error) {
	if len(elements) < 2 {
		return nil, errors.New("not enough elements to find max and min")
	}
	if perChange < 0 || perChange > 100 {
		return nil, errors.New("perChange must be between 0 and 100")
	}
	points := []float64{}
	points = append(points, elements[0].GetHighPrice())
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

	for _, e := range elements[1:] {
		if e.GetHighPrice() > points[len(points)-1] {
			add(e.GetOpenPrice(), up)
		} else {
			add(e.GetOpenPrice(), !up)
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
	g.process = append(g.process, ingest)
	go ingest.Start(ctx)
}

func (g *Gartley) nextKline(ctx context.Context, s string) (kline.Model, error) {
	ctx = context.Background()
	k, err := store.Kline.GetLast(ctx, s)
	if err != nil {
		return kline.Model{}, err
	}
	return k, nil
}
