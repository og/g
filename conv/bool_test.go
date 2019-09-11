package gconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func Test_BoolInt(t *testing.T)  {
	assert.Equal(t, int(1), BoolInt(true))
	assert.Equal(t, int(0), BoolInt(false))
	assert.Equal(t, int32(1), BoolInt32(true))
	assert.Equal(t, int32(0), BoolInt32(false))
	assert.Equal(t, int64(1), BoolInt64(true))
	assert.Equal(t, int64(0), BoolInt64(false))
	assert.Equal(t ,uint(1), BoolUint(true))
	assert.Equal(t ,uint(0), BoolUint(false))
	assert.Equal(t ,uint8(1), BoolUint8(true))
	assert.Equal(t ,uint8(0), BoolUint8(false))
	assert.Equal(t ,uint16(1), BoolUint16(true))
	assert.Equal(t ,uint16(0), BoolUint16(false))
	assert.Equal(t ,uint32(1), BoolUint32(true))
	assert.Equal(t ,uint32(0), BoolUint32(false))
	assert.Equal(t ,uint64(1), BoolUint64(true))
	assert.Equal(t ,uint64(0), BoolUint64(false))
}