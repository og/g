package gconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntBool(t *testing.T) {
	assert.Equal(t, true, IntBool(int(1)))
	assert.Equal(t, false, IntBool(int(0)))
}

func TestIntString(t *testing.T) {
	assert.Equal(t, "123456", IntString(int(123456)))
	assert.Equal(t, "123456", Int32String(int32(123456)))
	assert.Equal(t, "123456", Int64String(int64(123456)))
	assert.Equal(t, "11110001001000000", Int64StringWithBase(int64(123456), 2))
}
