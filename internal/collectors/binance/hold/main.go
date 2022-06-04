package main

import (
	"github.com/mcuv3/mcbot/internal/collectors/binance"
)

func main() {
	//	c := feed.NewClient()
	//	sym, err := c.ListSymbols(context.Background())
	binance.New().Start()
	//fmt.Println(sym, err)
}
