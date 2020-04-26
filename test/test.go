package gtest

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

type Assert struct {
	T *testing.T
}
func AS(t *testing.T) *Assert {
	return &Assert{T: t}
}
func (as Assert) Eql(expected, actual interface{}, msg ...interface{}) {
	assert.Equal(as.T,expected, actual, msg...)
}
func (as Assert) NoErrorSecond(v interface{}, err error) interface{} {
	as.NoError(err)
	return v
}
func (as Assert) HasErrorSecond(v interface{}, err error) interface{} {
	as.HasError(err)
	return v
}

func (as Assert) NoError(err error , msg ...interface{}) {
	assert.NoError(as.T, err, msg...)
}
func (as Assert) HasError( err error, msg ...interface{}) {
	assert.Error(as.T, err, msg...)
}
func (as Assert) EqualError(theError error, errString string, msg ...interface{}) {
	assert.EqualError(as.T, theError, errString, msg...)
}
func (as Assert) True(expression bool, msg ...interface{}) {
	assert.True(as.T, expression, msg...)
}
func (as Assert) False(expression bool, msg ...interface{}) {
	assert.False(as.T, expression, msg...)
}
func (as Assert) Len(sliceValue interface{}, len int, msg ...interface{}){
	assert.Len(as.T, sliceValue, len, msg...)
}
func (as Assert) Zero(i interface{}, msg ...interface{}){
	assert.Zero(as.T, i, msg...)
}
func (as Assert) RegexpString(regexpString string, str interface{}, msg ...interface{}){
	assert.Regexp(as.T, regexpString, str, msg...)
}
func (as Assert) Regexp(regexp *regexp.Regexp, str interface{}, msg ...interface{}){
	assert.Regexp(as.T, regexp, str, msg...)
}
func (as Assert) StringContains(fullString string, subString string, msg ...interface{}){
	assert.Contains(as.T, fullString, subString, msg...)
}
func (as Assert) StringSliceContains(stringSlice []string, subString string, msg ...interface{}){
	assert.Contains(as.T, stringSlice, subString, msg...)
}
func (as Assert) DirExists(path string, msg ...interface{}) {
	assert.DirExists(as.T, path, msg...)
}
// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// as.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func (as Assert) ElementsMatch(listA, listB interface{}, msg ...interface{}) (ok bool) {
	return assert.ElementsMatch(as.T, listA, listB, msg...)
}
func (as Assert) Range(v int, min int, max int, msg ...interface{}) {
	as.GtOrEql(v, min, msg...)
	assert.LessOrEqual(as.T, v, max, msg...)
}
func (as Assert) Gt(v int, expect int) {
	assert.Greater(as.T, v , expect)
}
func (as Assert) GtOrEql(v int, expect int, msg ...interface{}) {
	assert.GreaterOrEqual(as.T, v , expect, msg...)
}
func (as Assert) Lt(v int, expect int, msg ...interface{}) {
	assert.Less(as.T, v, expect, msg...)
}
func (as Assert) LtOrEql(v int, expect int, msg ...interface{}) {
	assert.LessOrEqual(as.T, v , expect, msg...)
}

// This is a very simple implementation, see the source code to better understand the role
func (as Assert) Run(n int, fn func(i int) (_break bool) ) {
	for i:=0; i<n; i++ {
		if fn(i) {
			break
		}
	}
}
