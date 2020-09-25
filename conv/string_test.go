package gconv

import (
	gtest "github.com/og/x/test"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStringInt(t *testing.T) {
	i , err := StringInt("123")
	assert.Equal(t, int(123), i)
	assert.Equal(t, nil, err)
}
func TestStringInt64(t *testing.T) {
	{
		i64 , err := StringInt64("123")
		assert.Equal(t, int64(123), i64)
		assert.Equal(t, nil, err)
	}
	{
		{
			i64 , err := StringInt64("-123")
			assert.Equal(t, int64(-123), i64)
			assert.Equal(t, nil, err)
		}
	}
}


func TestStringFloat64(t *testing.T) {
	i , err := StringFloat64("123.1")
	assert.Equal(t, float64(123.1), i)
	assert.Equal(t, nil, err)
}
func TestStringFloat32(t *testing.T) {
	i , err := StringFloat32("123.1")
	assert.Equal(t, float32(123.1), i)
	assert.Equal(t, nil, err)
}
func TestStringBool(t *testing.T) {
	{
		sList := []string{"true", "True","1", "t", "T"}
		for i:=0 ; i< len(sList) ; i++ {
			b, err := StringBool(sList[0])
			assert.Equal(t, b, true)
			assert.Equal(t, err, nil)
		}
	}
	{
		sList := []string{"false", "False","f", "F", "0"}
		for i:=0 ; i< len(sList) ; i++ {
			b, err := StringBool(sList[0])
			assert.Equal(t, b, false)
			assert.Equal(t, err, nil)
		}
	}
	{
		sList := []string{"asd", "","3t", "2e3f", "sd"}
		for i:=0 ; i< len(sList) ; i++ {
			b, err := StringBool(sList[0])
			assert.Equal(t, b, false)
			assert.EqualError(t, err, "og/x/conv: " + sList[0]  + " can't conv to bool")
		}
	}
}
func TestStringReflect(t *testing.T) {
	as := gtest.NewAS(t)
	type Data struct {
		s string
		i int
		i8 int8
		i16 int16
		i32 int32
		i64 int64
		ui uint
		ui8 uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		bool bool
		f32 float32
		f64 float64
		bytes []byte
	}
	data := Data{}
	as.NoError(
		StringReflect("s", reflect.ValueOf(&data.s)),
	)
	as.NoError(
		StringReflect("-1", reflect.ValueOf(&data.i)),
	)
	as.NoError(
		StringReflect("-2", reflect.ValueOf(&data.i8)),
	)
	as.NoError(
		StringReflect("-3", reflect.ValueOf(&data.i16)),
	)
	as.NoError(
		StringReflect("-4", reflect.ValueOf(&data.i32)),
	)
	as.NoError(
		StringReflect("-5", reflect.ValueOf(&data.i64)),
	)
	as.NoError(
		StringReflect("1", reflect.ValueOf(&data.ui)),
	)
	as.NoError(
		StringReflect("2", reflect.ValueOf(&data.ui8)),
	)
	as.NoError(
		StringReflect("3", reflect.ValueOf(&data.ui16)),
	)
	as.NoError(
		StringReflect("4", reflect.ValueOf(&data.ui32)),
	)
	as.NoError(
		StringReflect("5", reflect.ValueOf(&data.ui64)),
	)
	as.NoError(
		StringReflect("true", reflect.ValueOf(&data.bool)),
	)
	as.NoError(
		StringReflect("0.1", reflect.ValueOf(&data.f32)),
	)
	as.NoError(
		StringReflect("0.2", reflect.ValueOf(&data.f64)),
	)
	as.NoError(
		StringReflect("b我", reflect.ValueOf(&data.bytes)),
	)
	as.Equal(data, Data{"s", -1, -2, -3, -4, -5, 1, 2, 3, 4, 5, true, 0.1, 0.2, []byte("b我")})

}