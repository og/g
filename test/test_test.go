package gtest_test

import (
	"errors"
	json "github.com/og/x/json/lib"
	gtest "github.com/og/x/test"
	"log"
	"testing"
)

func TestAS(t *testing.T) {
	as := gtest.AS(t)
	// 不使用 NoErrorSecond
	{
		b , err := json.Marshal("a")
		as.NoError(err)
		as.Eql([]byte(`"a"`), b)
	}
	// 使用 NoErrorSecond
	{
		as.Eql([]byte(`"a"`), as.NoErrorSecond(json.Marshal("a")))
	}
	// 不使用 HasErrorSecond
	{
		b , err := json.Marshal(log.Print)
		as.HasError(err)
		as.Eql([]byte(nil), b)
	}
	// 使用 HasErrorSecond
	{
		// eql 是为了检查空值
		as.Eql([]byte(nil), as.HasErrorSecond(json.Marshal(log.Print)))
	}
	{
		as.Eql(errors.New("abc"), as.Panic(MockPanic))
		as.Eql("abc", as.Panic(MockPanicString))
	}
}

func MockPanic() {
	panic(errors.New("abc"))
}
func MockPanicString() {
	panic("abc")
}