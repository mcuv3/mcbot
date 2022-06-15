package stock

type Exchange = string

const (
	Binance Exchange = "Binance"
)

type Stock struct {
	Symbol     string   `bson:"symbol"`
	BaseAsset  string   `bson:"baseAsset"`
	QuoteAsset string   `bson:"quoteAsset"`
	Exchange   Exchange `bson:"exchange"`
}

var emptyStock = Stock{}

func (s Stock) CollectionName() string {
	return "stock"
}
