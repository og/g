package gtime

import (
	gtest "github.com/og/x/test"
	"testing"
)

func TestRange(t *testing.T) {
	as := gtest.NewAS(t)
	r := Range{}
	as.Equal(r.Type.Enum().Year.String(), "year")
	as.Equal(r.Type.Enum().Month.String(), "month")
	as.Equal(r.Type.Enum().Day.String(), "day")
	as.PanicError("RangeType: can not be (empty string)", func() {
		r.Type.Switch(
		func(_year int) {
			panic("can not be year")
		}, func(_month bool) {
			panic("can not be month")
		}, func(_day string) {
			panic("can not be day")
		})
	})
	{
		v := Range{}.Type.Enum().Year
		exec := false
		v.Switch(
		func(_year int) {
			exec = true
		}, func(_month bool) {
			panic("can not be month")
		}, func(_day string) {
			panic("can not be day")
		})
		as.True(exec)
	}
	{
		exec := false
		v := Range{}.Type.Enum().Day
		v.Switch(
			func(_year int) {
				panic("can not be year")
			}, func(_month bool) {
				panic("can not be month")
			}, func(_day string) {
				exec = true
			})
			as.True(exec)
	}
	{
		exec := false
		v := Range{}.Type.Enum().Month
		v.Switch(
			func(_year int) {
				panic("can not be year")
			}, func(_month bool) {
				exec = true
			}, func(_day string) {
				panic("can not be day")
			})
			as.True(exec)
	}
}
