package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2022-03-02", quote.Hour2, true)
	fmt.Print(spy.CSV())
	fmt.Println(spy.Volume)
}
