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
