package stock

type Exchange = string

const (
	Binance Exchange = "Binance"
)

type Stock struct {
	ID       string   `bson:"_id"`
	Symbol   string   `bson:"symbol"`
	Name     string   `bson:"name"`
	Exchange Exchange `bson:"exchange"`
}

var emptyStock = Stock{}

func (s Stock) CollectionName() string {
	return "stock"
}
