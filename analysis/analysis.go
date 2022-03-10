package analysis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/MauricioAntonioMartinez/mcbot/shared"
)

type Analyzer struct {
	frame       int
	symbol      string
	trendSize   shared.TrendSize
	cryptostore shared.Collector
}

func NewAnalyser(params shared.AnalyserParams) *Analyzer {
	return &Analyzer{
		frame:       params.Frame,
		symbol:      params.Symbol,
		trendSize:   params.TrendSize,
		cryptostore: params.Collector,
	}
}

func (a *Analyzer) Analyse() ([]shared.Trend, error) {
	data, err := a.cryptostore.Collect(a.symbol, a.frame)
	if err != nil {
		return nil, err
	}

	//	for _, v := range data {

	//fmt.Printf("%s v: %v h:%v l:%v time: %v \n", a.symbol, v.Close, v.High, v.Low, a.getDate(v.Time))
	//}

	fmt.Println(len(data))
	return nil, nil
}

func (a *Analyzer) getDate(ux string) time.Time {
	i, err := strconv.ParseInt(ux, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}
