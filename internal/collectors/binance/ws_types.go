package binance

type Methods string

const (
	SUBSCRIBE   Methods = "SUBSCRIBE"
	UNSUBSCRIBE Methods = "UNSUBSCRIBE"
)

type Payloads interface {
	AggregateTradePayload | KlinePayload
}

type AggregateTradePayload struct {
	Event            string `json:"e"`
	EventTime        int    `json:"E"`
	Symbol           string `json:"s"`
	AggregateTradeID int    `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int    `json:"f"`
	LastTradeID      int    `json:"l"`
	TradeTime        int    `json:"T"`
	IsBuyerMarket    bool   `json:"m"`
	Ignore           bool   `json:"M"`
}

type KlinePayload struct {
	Event     string `json:"e"`
	EventTime int    `json:"E"`
	Symbol    string `json:"s"`
	K         struct {
		StartTime    int     `json:"t"`
		CloseTime    int     `json:"T"`
		Symbol       string  `json:"s"`
		Interval     string  `json:"i"`
		FirstTradeID int     `json:"f"`
		LastTradeID  int     `json:"L"`
		OpenPrice    string  `json:"o"`
		ClosePrice   string  `json:"c"`
		HighPrice    float64 `json:"h"`
		LowPrice     string  `json:"l"`
		Volume       string  `json:"v"`
		NTrades      int     `json:"n"`
		IsClosed     bool    `json:"x"`
		QuoteVolume  string  `json:"q"`
	} `json:"k"`
}
