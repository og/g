package gtest_test

import (
	"errors"
	json "github.com/og/x/json/lib"
	gtest "github.com/og/x/test"
	"log"
	"testing"
)

func TestAS(t *testing.T) {
	as := gtest.NewAS(t)
	// 不使用 NoErrorSecond
	{
		b , err := json.Marshal("a")
		as.NoError(err)
		as.Equal([]byte(`"a"`), b)
	}
	// 使用 NoErrorSecond
	{
		as.Equal([]byte(`"a"`), as.NoErrorSecond(json.Marshal("a")))
	}
	// 不使用 HasErrorSecond
	{
		b , err := json.Marshal(log.Print)
		as.HasError(err)
		as.Equal([]byte(nil), b)
	}
	// 使用 HasErrorSecond
	{
		// eql 是为了检查空值
		as.Equal([]byte(nil), as.HasErrorSecond(json.Marshal(log.Print)))
	}
	{
		as.Equal(errors.New("abc"), as.Panic(MockPanic))
		as.Equal("abc", as.Panic(MockPanicString))
		as.PanicError(errors.New("abc"), MockPanic)
		as.PanicErrorString("abc", MockPanic)
	}
}

func MockPanic() {
	panic(errors.New("abc"))
}
func MockPanicString() {
	panic("abc")
}