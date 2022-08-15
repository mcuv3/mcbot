package kline

import (
	"log"
	"strconv"
)

var collectionName = "kline"

type Model struct {
	Symbol     string `bson:"symbol"`
	StartTime  int    `bson:"start_time"`
	CloseTime  int    `bson:"close_time"`
	OpenPrice  string `bson:"open_price"`
	ClosePrice string `bson:"close_price"`
	HighPrice  string `bson:"high_price"`
	LowPrice   string `bson:"low_price"`
	CreatedAt  int64  `bson:"created_at"`
}

var emptyTrend = Model{}

func (t Model) Table() string {
	return collectionName
}

func (t Model) GetHighPrice() float64 {
	return toFloat(t.HighPrice)
}

func (t Model) GetLowPrice() float64 {
	return toFloat(t.LowPrice)
}

func (t Model) GetOpenPrice() float64 {
	return toFloat(t.OpenPrice)
}

func (t Model) GetClosePrice() float64 {
	return toFloat(t.ClosePrice)
}

func (t Model) GetDetails() [4]float64 {
	return [4]float64{
		t.GetOpenPrice(),
		t.GetClosePrice(),
		t.GetLowPrice(),
		t.GetHighPrice(),
	}
}

func toFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("Error converting string to float:", err)
		return 0
	}
	return f
}
