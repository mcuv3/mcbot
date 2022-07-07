package kline

var collectionName = "kline"

type Model struct {
	Symbol     string `bson:"symbol"`
	StartTime  int    `bson:"start_time"`
	CloseTime  int    `bson:"close_time"`
	OpenPrice  string `bson:"open_price"`
	ClosePrice string `bson:"close_price"`
	HighPrice  string `bson:"high_price"`
	LowPrice   string `bson:"low_price"`
}

var emptyTrend = Model{}

func (t Model) Table() string {
	return collectionName
}
