package glist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyList_Push(t *testing.T) {
	aList := AnyList{}
	aList.Push("name", "nimo")
	assert.Equal(t, []interface{}{"name","nimo"}, aList.Value)
	aList.Push("age", 18)
	assert.Equal(t, []interface{}{"name","nimo","age", 18}, aList.Value)
}
