package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciRetro(t *testing.T) {
	t.Run("dow trend", func(t *testing.T) {
		l := GetFibonacciRetrace(21700, 21600)
		assert.Equal(t, 21623.6, l.L236)
		assert.Equal(t, 21638.2, l.L382)

	})

	t.Run("up trend", func(t *testing.T) {
		l := GetFibonacciRetrace(21600, 21700)
		assert.Equal(t, 21638.2, l.L618)
		assert.Equal(t, 21661.8, l.L382)

	})
}
