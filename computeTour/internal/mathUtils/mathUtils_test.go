package mathUtils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {

	assert.Equal(t, Factorial(0), 1, "!0 should equal 1")
	assert.Equal(t, Factorial(1), 1, "!1 should equal 1")
	assert.Equal(t, Factorial(2), 2, "!2 should equal 2")
	assert.Equal(t, Factorial(3), 6, "!3 should equal 6")
	assert.Equal(t, Factorial(4), 24, "!4 should equal 24")
	assert.Equal(t, Factorial(5), 120, "!5 should equal 120")
	assert.Equal(t, Factorial(10), 3628800, "!10 should equal 3 628 800")
}

func TestDistance(t *testing.T) {
	lat1 := 48.87419826061369
	long1 := 2.3281629241227764
	lat2 := 48.87288862207115
	long2 := 2.331235505332878

	assert.Equal(t, math.Floor(Distance(lat1, long1, lat2, long2, "m")), 267.0, "Distance between those two coordinates should be 267 meters")
}
func TestRawDistance(t *testing.T) {
	lat1 := 48.87419826061369
	long1 := 2.3281629241227764
	lat2 := 48.87288862207115
	long2 := 2.331235505332878

	assert.Equal(t, RawDistance(lat1, long1, lat2, long2), 4.2030272437543204e-05, "Distance between those two coordinates should be 0.9999999991167281")
}
