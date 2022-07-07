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
	Level1 float64
	Level2 float64
	Level3 float64
	Level4 float64
	Level5 float64
	Level6 float64
	Level7 float64
	Level8 float64
}

func GetUpTrendFibonacci(h, l float64) RetracementLevels {
	return RetracementLevels{
		Level1: h - (h-l)*0.236,
		Level2: h - (h-l)*0.382,
		Level3: h - (h-l)*0.500,
		Level4: h - (h-l)*0.618,
		Level5: h - (h-l)*0.764,
		Level6: h - (h-l)*1.00,
		Level7: h - (h-l)*1.382,
		Level8: h - (h-l)*1.618,
	}
}

func GetDownTrendFibonacci(h, l float64) RetracementLevels {
	return RetracementLevels{
		Level1: l + (h-l)*0.236,
		Level2: l + (h-l)*0.382,
		Level3: l + (h-l)*0.500,
		Level4: l + (h-l)*0.618,
		Level5: l + (h-l)*0.764,
		Level6: l + (h-l)*1.00,
		Level7: l + (h-l)*1.382,
		Level8: l + (h-l)*1.618,
	}
}
