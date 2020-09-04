package gmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntPercent(t *testing.T) {
	assert.Equal(t, 33, IntPercent(333,1000))
	assert.Equal(t, 34, IntPercent(344,1000))
	assert.Equal(t, 34, IntPercent(345,1000))
	assert.Equal(t, 35, IntPercent(346,1000))
	assert.Equal(t, 35, IntPercent(349,1000))
	assert.Equal(t, 40, IntPercent(4,10))
	assert.Equal(t, 67, IntPercent(41,61))
	assert.Equal(t, 0, IntPercent(41,0))
	assert.Equal(t, 0, IntPercent(0,41))
	assert.Equal(t, 2000, IntPercent(100,5))
}

func TestFloat64Percent(t *testing.T) {
	assert.Equal(t, int(33), Float64Percent(0.333,1))
	assert.Equal(t, 34, Float64Percent(0.344,1))
	assert.Equal(t, 34, Float64Percent(0.345,1))
	assert.Equal(t, 35, Float64Percent(0.346,1))
	assert.Equal(t, 35, Float64Percent(0.349,1))
	assert.Equal(t, int(20), Float64Percent(0.4,2))
	assert.Equal(t, int(67), Float64Percent(0.41,0.61))
	assert.Equal(t, int(0), Float64Percent(41,0))
	assert.Equal(t, int(0), Float64Percent(0,41))
	assert.Equal(t, int(2000), Float64Percent(100,5))
}


func TestFloat64ToFixed (t *testing.T) {
	v := float64(0.123456789123456789123456789)
	v1 := float64(1.123456789123456789123456789)
	assert.Equal(t, Float64ToFixed(v1, 0), float64(1))
	assert.Equal(t, Float64ToFixed(v, 0), float64(0))
	assert.Equal(t, Float64ToFixed(v, 1), float64(0.1))
	assert.Equal(t, Float64ToFixed(v, 2), float64(0.12))
	assert.Equal(t, Float64ToFixed(v, 3), float64(0.123))
	assert.Equal(t, Float64ToFixed(v, 4), float64(0.1235))
	assert.Equal(t, Float64ToFixed(v, 5), float64(0.12346))
	assert.Equal(t, Float64ToFixed(v, 6), float64(0.123457))
	assert.Equal(t, Float64ToFixed(v, 7), float64(0.1234568))
	assert.Equal(t, Float64ToFixed(v, 8), float64(0.12345679))
	assert.Equal(t, Float64ToFixed(v, 9), float64(0.123456789))
	assert.Equal(t, Float64ToFixed(v, 10), float64(0.1234567891))
}
