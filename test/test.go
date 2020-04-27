package gtest

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

type AS struct {
	T *testing.T
}
func NewAS(t *testing.T) *AS {
	return &AS{T: t}
}
func (as AS) Equal(expected, actual interface{}, msg ...interface{}) {
	assert.Equal(as.T,expected, actual, msg...)
}
func (as AS) NoErrorSecond(v interface{}, err error) interface{} {
	as.NoError(err)
	return v
}
func (as AS) HasErrorSecond(v interface{}, err error) interface{} {
	as.HasError(err)
	return v
}

func (as AS) NoError(err error , msg ...interface{}) {
	assert.NoError(as.T, err, msg...)
}
func (as AS) HasError( err error, msg ...interface{}) {
	assert.Error(as.T, err, msg...)
}
func (as AS) ErrorString(theError error, errString string, msg ...interface{}) {
	assert.EqualError(as.T, theError, errString, msg...)
}
func (as AS) Error(theError error, err error, msg ...interface{}) {
	assert.EqualError(as.T, theError, err.Error(), msg...)
}
func (as AS) True(expression bool, msg ...interface{}) {
	assert.True(as.T, expression, msg...)
}
func (as AS) False(expression bool, msg ...interface{}) {
	assert.False(as.T, expression, msg...)
}
func (as AS) Len(sliceValue interface{}, len int, msg ...interface{}){
	assert.Len(as.T, sliceValue, len, msg...)
}
func (as AS) Zero(i interface{}, msg ...interface{}){
	assert.Zero(as.T, i, msg...)
}
func (as AS) RegexpString(regexpString string, str interface{}, msg ...interface{}){
	assert.Regexp(as.T, regexpString, str, msg...)
}
func (as AS) Regexp(regexp *regexp.Regexp, str interface{}, msg ...interface{}){
	assert.Regexp(as.T, regexp, str, msg...)
}
func (as AS) StringContains(fullString string, subString string, msg ...interface{}){
	assert.Contains(as.T, fullString, subString, msg...)
}
func (as AS) StringSliceContains(stringSlice []string, subString string, msg ...interface{}){
	assert.Contains(as.T, stringSlice, subString, msg...)
}
func (as AS) DirExists(path string, msg ...interface{}) {
	assert.DirExists(as.T, path, msg...)
}
// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// as.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func (as AS) ElementsMatch(listA, listB interface{}, msg ...interface{}) (ok bool) {
	return assert.ElementsMatch(as.T, listA, listB, msg...)
}
func (as AS) Range(v int, min int, max int, msg ...interface{}) {
	as.GtOrEql(v, min, msg...)
	assert.LessOrEqual(as.T, v, max, msg...)
}
func (as AS) Gt(v int, expect int) {
	assert.Greater(as.T, v , expect)
}
func (as AS) GtOrEql(v int, expect int, msg ...interface{}) {
	assert.GreaterOrEqual(as.T, v , expect, msg...)
}
func (as AS) Lt(v int, expect int, msg ...interface{}) {
	assert.Less(as.T, v, expect, msg...)
}
func (as AS) LtOrEql(v int, expect int, msg ...interface{}) {
	assert.LessOrEqual(as.T, v , expect, msg...)
}

// This is a very simple implementation, see the source code to better understand the role
func (as AS) Run(n int, fn func(i int) (_break bool) ) {
	for i:=0; i<n; i++ {
		if fn(i) {
			break
		}
	}
}
func (as AS) Nil(v interface{}, msg ...interface{}) {
	assert.Nil(as.T, v, msg...)
}
func (as AS) NotNil(v interface{}, msg ...interface{}) {
	assert.NotNil(as.T, v, msg...)
}
func (as AS) Panic(fn func ()) (recoverValue interface{})  {
	defer func() {
		r := recover()
		as.NotNil(r)
		recoverValue = r
	}()
	fn()
	return
}
func (as AS) PanicErrorString(expectErrorString string, fn func ())  {
	as.Equal(errors.New(expectErrorString), as.Panic(fn))
}
func (as AS) PanicError(expectError error, fn func ())  {
	as.Equal(expectError, as.Panic(fn))
}