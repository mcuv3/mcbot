package stock

import "github.com/google/uuid"

type Exchange = string

const (
	Binance Exchange = "Binance"
)

type Stock struct {
	ID       uuid.UUID `bson:"id"`
	Symbol   string    `bson:"symbol"`
	Name     string    `bson:"name"`
	Exchange Exchange  `bson:"exchange"`
}

var emptyStock = Stock{}

func (s Stock) Table() string {
	return "stocks"
}
