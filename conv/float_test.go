package gconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64String(t *testing.T) {
	assert.Equal(t, Float64String(float64(1.222222222222222222222222222222)), "1.2222222222222223")
	assert.Equal(t, Float64String(float64(1.222)), "1.222")
	assert.Equal(t, Float64String(float64(.222)), "0.222")
}
func TestFloat32String(t *testing.T) {
	assert.Equal(t, Float32String(float32(1.222222222222222222222222222222)), "1.2222222")
	assert.Equal(t, Float32String(float32(1.222)), "1.222")
	assert.Equal(t, Float32String(float32(.222)), "0.222")
}

