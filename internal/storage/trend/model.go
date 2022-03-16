package trend

import (
	"time"

	"github.com/google/uuid"
)

// Amount of horus of a given trend
type TrendSize = int

type Trend struct {
	ID           uuid.UUID `bson:"id"`
	Avg          float64   `bson:"average"`
	Volatility   float64   `bson:"volatility"`
	Size         TrendSize `bson:"size"`
	AvgVolume    float64   `bson:"avg_volume"`
	StdDeviation float64   `bson:"std_deviation"`
	Start        time.Time `bson:"start"`
	End          time.Time `bson:"end"`
}

var emptyTrend = Trend{}

func (t Trend) Table() string {
	return "trends"
}
