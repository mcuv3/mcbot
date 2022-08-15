package logic

import (
	"context"
	"log"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mcuv3/mcbot/internal/ingestors"
	"github.com/mcuv3/mcbot/internal/shared"
)

func (l Logic) GraphAnalysis(ctx context.Context, params shared.GraphAnalysisParams) (*charts.Kline, error) {
	kline := charts.NewKLine()

	klines, err := l.Kline.List(ctx, map[string]interface{}{
		"symbol": params.Symbol,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(len(klines))

	res, err := ingestors.FindMaxAndMinPoints(klines[70:], 0.060)
	log.Println(res, err)

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for _, k := range klines {
		x = append(x, time.Unix(int64(k.CreatedAt), 0).Format(time.RFC822))
		y = append(y, opts.KlineData{Value: k.GetDetails()})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Graph Analysis",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	kline.SetXAxis(x).AddSeries("kline", y).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(opts.MarkPointNameTypeItem{
				Name:     "highest value",
				Type:     "max",
				ValueDim: "highest",
			}),
			charts.WithMarkPointNameTypeItemOpts(opts.MarkPointNameTypeItem{
				Name:     "lowest value",
				Type:     "min",
				ValueDim: "lowest",
			}),
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show: true,
				},
			}),
		)

	return kline, nil
}
