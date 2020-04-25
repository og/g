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
func (as Assert) Eql(expected, actual interface{}, msgAndArgs ...interface{}) {
	assert.Equal(as.T,expected, actual, msgAndArgs...)
}
func (as Assert) NoErrorSecond(v interface{}, err error) interface{} {
	as.NoError(err)
	return v
}
func (as Assert) HasErrorSecond(v interface{}, err error) interface{} {
	as.HasError(err)
	return v
}

func (as Assert) NoError(err error , msgAndArgs ...interface{}) {
	assert.NoError(as.T, err, msgAndArgs...)
}
func (as Assert) HasError( err error, msgAndArgs ...interface{}) {
	assert.Error(as.T, err, msgAndArgs...)
}
func (as Assert) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	assert.EqualError(as.T, theError, errString, msgAndArgs...)
}
func (as Assert) True(expression bool, msgAndArgs ...interface{}) {
	assert.True(as.T, expression, msgAndArgs...)
}
func (as Assert) False(expression bool, msgAndArgs ...interface{}) {
	assert.False(as.T, expression, msgAndArgs...)
}
func (as Assert) Len(sliceValue interface{}, len int, msgAndArgs ...interface{}){
	assert.Len(as.T, sliceValue, len, msgAndArgs...)
}
func (as Assert) Zero(i interface{}, msgAndArgs ...interface{}){
	assert.Zero(as.T, i, msgAndArgs...)
}
func (as Assert) RegexpString(regexpString string, str interface{}, msgAndArgs ...interface{}){
	assert.Regexp(as.T, regexpString, str, msgAndArgs...)
}
func (as Assert) Regexp(regexp *regexp.Regexp, str interface{}, msgAndArgs ...interface{}){
	assert.Regexp(as.T, regexp, str, msgAndArgs...)
}
func (as Assert) StringContains(fullString string, subString string, msgAndArgs ...interface{}){
	assert.Contains(as.T, fullString, subString, msgAndArgs...)
}
func (as Assert) StringSliceContains(stringSlice []string, subString string, msgAndArgs ...interface{}){
	assert.Contains(as.T, stringSlice, subString, msgAndArgs...)
}
func (as Assert) DirExists(path string, msgAndArgs ...interface{}) {
	assert.DirExists(as.T, path, msgAndArgs...)
}
// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// as.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func (as Assert) ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	return assert.ElementsMatch(as.T, listA, listB, msgAndArgs...)
}