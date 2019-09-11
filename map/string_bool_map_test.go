package gmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringBool(t *testing.T) {
	sbMap := NewStringBool()
	{
		value, has := sbMap.Get("nimo")
		assert.Equal(t, false, value)
		assert.Equal(t, false, has)
	}
	{
		sbMap.Set("nimo", true)
		value, has := sbMap.Get("nimo")
		assert.Equal(t, true, value)
		assert.Equal(t, true, has)
	}
	{
		assert.Equal(t, true, sbMap.Has("nimo"))
		assert.Equal(t, false, sbMap.Has("nico"))
		sbMap.Set("nico", true)
		assert.Equal(t, true, sbMap.Has("nimo"))
	}
	{
		sbMap.Remove("nimo")
		assert.Equal(t, false, sbMap.Has("nimo"))
		value, has := sbMap.Get("nimo")
		assert.Equal(t, false, value)
		assert.Equal(t, false, has)
	}
	{
		assert.Equal(t, map[string]bool(map[string]bool{"nico":true}), sbMap.Value())
	}
	{
		assert.Equal(t, 1, sbMap.Size())
	}
	{
		sbMap.Clear()
		assert.Equal(t, false, sbMap.Has("nimo"))
		assert.Equal(t, false, sbMap.Has("nico"))
		assert.Equal(t, map[string]bool(map[string]bool{}), sbMap.Value())
	}
}