package shared

import (
	"math"
	"strconv"
	"time"
)

// DivUp returns the least integer value greater than or equal to the quotient a/b.
func DivUp(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

func UnixToDate(ux string) time.Time {
	i, err := strconv.ParseInt(ux, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}

//0.236, 0.382, 0.500, 0.618, 0.764, 1.00, 1.382, 1.618

type RetracementLevels struct {
	L236  float64
	L382  float64
	L500  float64
	L618  float64
	L786  float64
	L100  float64
	L1382 float64
	L1618 float64
}

func GetUpTrendFibonacci(h, l float64) RetracementLevels {
	return RetracementLevels{
		L236:  h - (h-l)*0.236,
		L382:  h - (h-l)*0.382,
		L500:  h - (h-l)*0.500,
		L618:  h - (h-l)*0.618,
		L786:  h - (h-l)*0.786,
		L100:  h - (h-l)*1.00,
		L1382: h - (h-l)*1.382,
		L1618: h - (h-l)*1.618,
	}
}

func GetDownTrendFibonacci(h, l float64) RetracementLevels {
	return RetracementLevels{
		L236:  l + (h-l)*0.236,
		L382:  l + (h-l)*0.382,
		L500:  l + (h-l)*0.500,
		L618:  l + (h-l)*0.618,
		L786:  l + (h-l)*0.786,
		L100:  l + (h-l)*1.00,
		L1382: l + (h-l)*1.382,
		L1618: l + (h-l)*1.618,
	}
}
