package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciRetro(t *testing.T) {
	t.Run("up trend", func(t *testing.T) {
		l := GetUpTrendFibonacci(1000, 800)
		assert.Equal(t, 952.8, l.L236)
		assert.Equal(t, 923.6, l.L382)

	})
}
