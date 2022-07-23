package ingestors

import (
	"testing"

	"github.com/mcuv3/mcbot/internal/storage/kline"
	"github.com/stretchr/testify/assert"
)

func TestSelectInfexionPoints(t *testing.T) {
	var upTrend = []kline.Model{
		{
			HighPrice: 1,
		},
		{
			HighPrice: 2,
		},
		{
			HighPrice: 3,
		},
		{
			HighPrice: 4,
		},
		{
			HighPrice: 5,
		},
		{
			HighPrice: 6,
		},
		{
			HighPrice: 4,
		},
		{
			HighPrice: 1,
		},
		{
			HighPrice: 5,
		},
		{
			HighPrice: 8,
		},
		{
			HighPrice: 10,
		},
		{
			HighPrice: 3,
		},
	}

	var downTrend = []kline.Model{
		{
			HighPrice: 10,
		},
		{
			HighPrice: 7,
		},
		{
			HighPrice: 2,
		},
		{
			HighPrice: 8,
		},
		{
			HighPrice: 10,
		},
		{
			HighPrice: 6,
		},
		{
			HighPrice: 2,
		},
		{
			HighPrice: 1,
		},
		{
			HighPrice: 5,
		},
		{
			HighPrice: 8,
		},
		{
			HighPrice: 9,
		},
		{
			HighPrice: 10,
		},
	}

	t.Run("up trend", func(t *testing.T) {
		elements, err := findMaxAndMinPoints(upTrend, 25)
		assert.NoError(t, err)
		assert.Exactly(t, []float64{1, 6, 1, 10, 3}, elements)
	})
	t.Run("down trend", func(t *testing.T) {
		elements, err := findMaxAndMinPoints(downTrend, 25)
		assert.NoError(t, err)
		assert.Exactly(t, []float64{10, 2, 10, 1, 10}, elements)
	})

}
