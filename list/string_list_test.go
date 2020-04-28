package glist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringList_Concat (t *testing.T) {
	sList := StringList{}
	sList.Concat([]string{"a"})
	assert.Equal(t, []string{"a"}, sList.Value)
	sList.Concat([]string{"b", "c"})
	assert.Equal(t, []string{"a","b","c"}, sList.Value)
}

func TestStringList_Push(t *testing.T) {
	sList := StringList{}
	sList.Push("a")
	assert.Equal(t, []string{"a"}, sList.Value)
	sList.Push("b", "c")
	assert.Equal(t, []string{"a","b","c"}, sList.Value)

	otherStringList := StringList{}
	otherStringList.Push("abc")
	assert.Equal(t, []string{"abc"}, otherStringList.Value)
}

func TestStringList_Pop(t *testing.T) {
	sList := StringList{[]string{"a","b"}}
	sList.Pop()
	assert.Equal(t, []string{"a"}, sList.Value)
	sList.Pop()
	assert.Equal(t, []string{}, sList.Value)
	sList.Pop() // When len(sList.Value) == 0 , will not panic error
	assert.Equal(t, []string{}, sList.Value)
}

func TestStringList_PopBind(t *testing.T) {
	sList := StringList{[]string{"a","b"}}
	lastString := StringListBindValue{}
	sList.PopBind(&lastString)
	assert.Equal(t, []string{"a"}, sList.Value)
	assert.Equal(t, StringListBindValue{"b", true}, lastString)
	sList.PopBind(&lastString)
	assert.Equal(t, []string{}, sList.Value)
	assert.Equal(t, StringListBindValue{"a", true}, lastString)
	sList.PopBind(&lastString)
	assert.Equal(t, []string{}, sList.Value)
	assert.Equal(t, StringListBindValue{"", false}, lastString)
}

func TestStringList_Shift(t *testing.T) {
	sList := StringList{[]string{"a","b"}}
	sList.Shift()
	assert.Equal(t, []string{"b"}, sList.Value)
	sList.Shift()
	assert.Equal(t, []string{}, sList.Value)
	sList.Shift() // When len(sList.Value) == 0 , will not panic error
	assert.Equal(t, []string{}, sList.Value)
}

func TestStringList_ShiftBind(t *testing.T) {
	sList := StringList{[]string{"a","b"}}
	var firstString StringListBindValue
	sList.ShiftBind(&firstString)
	assert.Equal(t, StringListBindValue{"a", true}, firstString)
	assert.Equal(t, []string{"b"}, sList.Value)

	sList.ShiftBind(&firstString)
	assert.Equal(t, StringListBindValue{"b", true}, firstString)
	assert.Equal(t, []string{}, sList.Value)

	sList.ShiftBind(&firstString) // When len(sList.Value) == 0 , will not panic error
	assert.Equal(t, []string{}, sList.Value)
	assert.Equal(t, StringListBindValue{"", false}, firstString)

}

func TestStringList_Unshift(t *testing.T) {
	var sList StringList
	sList.Unshift("a")
	assert.Equal(t, []string{"a"}, sList.Value)
	sList.Unshift("b")
	assert.Equal(t, []string{"b", "a"}, sList.Value)
}
func TestStringList_Find(t *testing.T) {
	{
		sList := StringList{[]string{"a","b","c"}}
		var eachIndexList []int
		resultIndex, result := sList.Find(func(index int, item string) bool {
			eachIndexList = append(eachIndexList, index)
			return item == "b"
		})
		assert.Equal(t, 1, resultIndex)
		assert.Equal(t, true, result)
		assert.Equal(t, []int{0,1}, eachIndexList)
	}
	{
		sList := StringList{[]string{"a","b","c"}}
		var eachIndexList []int
		resultIndex, result := sList.Find(func(index int, item string) bool {
			eachIndexList = append(eachIndexList, index)
			return item == "z"
		})
		assert.Equal(t, -1, resultIndex)
		assert.Equal(t, false, result)
		assert.Equal(t, []int{0,1,2}, eachIndexList)
	}
}

func TestStringList_CheckAll(t *testing.T) {
	{
		sList := StringList{[]string{"a","b","c"}}
		var eachIndexList []int
		result := sList.CheckAll(func(index int, item string) bool {
			eachIndexList = append(eachIndexList, index)
			return len(item) == 1
		})
		assert.Equal(t, true, result)
		assert.Equal(t, []int{0,1,2}, eachIndexList)
	}
	{
		sList := StringList{[]string{"a","bb","c"}}
		var eachIndexList []int
		result := sList.CheckAll(func(index int, item string) bool {
			eachIndexList = append(eachIndexList, index)
			return len(item) == 1
		})
		assert.Equal(t, false, result)
		assert.Equal(t, []int{0,1}, eachIndexList)
	}
}

func TestStringList_Join(t *testing.T) {
	{
		sList := StringList{[]string{"a","b","c"}}

		assert.Equal(t, "a-b-c", sList.Join("-"))
	}
}
func TestStringList_In(t *testing.T) {
	{
		sList := StringList{[]string{"a","b","c"}}
		assert.Equal(t, sList.In("a"), true)
		assert.Equal(t, sList.In("b"), true)
		assert.Equal(t, sList.In("c"), true)
		assert.Equal(t, sList.In("d"), false)
	}
}