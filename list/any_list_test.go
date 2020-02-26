package glist

import (
	"github.com/matryer/is"
	"testing"
)

func TestAnyList_Push(t *testing.T) {
	is := is.New(t)

	aList := AnyList{}
	aList.Push("name", "nimo")
	is.Equal([]interface{}{"name","nimo"}, aList.Value)
	aList.Push("age", 18)
	is.Equal([]interface{}{"name","nimo","age", 18}, aList.Value)
}
