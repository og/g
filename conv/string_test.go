package gconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringInt(t *testing.T) {
	i , err := StringInt("123")
	assert.Equal(t, int(123), i)
	assert.Equal(t, nil, err)
}
func TestStringInt64(t *testing.T) {
	i64 , err := StringInt64("123")
	assert.Equal(t, int64(123), i64)
	assert.Equal(t, nil, err)
}


func TestStringFloat64(t *testing.T) {
	i , err := StringFloat64("123.1")
	assert.Equal(t, float64(123.1), i)
	assert.Equal(t, nil, err)
}
func TestStringFloat32(t *testing.T) {
	i , err := StringFloat32("123.1")
	assert.Equal(t, float32(123.1), i)
	assert.Equal(t, nil, err)
}