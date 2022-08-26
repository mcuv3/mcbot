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

// up - add to low
// down - reduce it from high

func GetFibonacciRetrace(a, b float64) RetracementLevels {
	base := b
	var flip float64 = -1
	if a > b { // down trend
		base = b
		flip = 1
	}

	return RetracementLevels{
		L236:  base + flip*math.Abs((a-b)*0.236),
		L382:  base + flip*math.Abs((a-b)*0.382),
		L500:  base + flip*math.Abs((a-b)*0.500),
		L618:  base + flip*math.Abs((a-b)*0.618),
		L786:  base + flip*math.Abs((a-b)*0.786),
		L100:  base + flip*math.Abs((a-b)*1.00),
		L1382: base + flip*math.Abs((a-b)*1.382),
		L1618: base + flip*math.Abs((a-b)*1.618),
	}
}
