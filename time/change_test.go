package gtime_test

import (
	gtest "github.com/og/x/test"
	gtime "github.com/og/x/time"
	"testing"
)

func TestFirstNanoseconds(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	layout := "2006-01-02 15:04:05.999999999"
	{
		testCase := [][]string{
			{"1992-12-19 11:11:00.763035", "1992-12-19 11:11:00",},
			{"1992-12-19 11:11:01.763035", "1992-12-19 11:11:01",},
			{"1992-12-19 11:11:02.763035", "1992-12-19 11:11:02",},
		}
		for _,v := range testCase {
			tv := gtime.ParseUTC(layout, v[0])
			as.Equal(gtime.FirstNanoseconds(tv).Format(layout), v[1])
			as.Equal(gtime.FirstNanoseconds(tv).Nanosecond(), 0)
		}
	}
}
func TestFirstSecond(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-19 11:11:11", "1992-12-19 11:11:00",},
			{"1992-12-19 11:11:00", "1992-12-19 11:11:00",},
			{"1992-12-19 11:11:01", "1992-12-19 11:11:00",},
			{"1992-12-19 11:11:59", "1992-12-19 11:11:00",},
			{"1992-12-19 11:11:58", "1992-12-19 11:11:00",},
		}
		for _,v := range testCase {
			tv := gtime.ParseUTC(gtime.LayoutTime, v[0])
			as.Equal(gtime.FirstSecond(tv).Format(gtime.LayoutTime), v[1])
			as.Equal(tv.Nanosecond(), 0)
		}
	}
}

func TestLastSecond(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-19 11:11:11", "1992-12-19 11:11:59",},
			{"1992-12-19 11:11:00", "1992-12-19 11:11:59",},
			{"1992-12-19 11:11:01", "1992-12-19 11:11:59",},
			{"1992-12-19 11:11:59", "1992-12-19 11:11:59",},
			{"1992-12-19 11:11:58", "1992-12-19 11:11:59",},
		}
		for _,v := range testCase {
			tv := gtime.ParseUTC(gtime.LayoutTime, v[0])
			as.Equal(gtime.LastSecond(tv).Format(gtime.LayoutTime), v[1])
			as.Equal(tv.Nanosecond(), 0)
		}
	}
}

func TestFirstMinute(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-19 11:00:11", "1992-12-19 11:00:00",},
			{"1992-12-19 11:01:11", "1992-12-19 11:00:00",},
			{"1992-12-19 11:02:11", "1992-12-19 11:00:00",},
			{"1992-12-19 11:58:11", "1992-12-19 11:00:00",},
			{"1992-12-19 11:59:11", "1992-12-19 11:00:00",},
		}
		for _,v := range testCase {
			as.Equal(gtime.FirstMinute(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}

func TestLastMinute(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-19 11:01:11", "1992-12-19 11:59:59",},
			{"1992-12-19 11:02:11", "1992-12-19 11:59:59",},
			{"1992-12-19 11:03:11", "1992-12-19 11:59:59",},
			{"1992-12-19 11:58:11", "1992-12-19 11:59:59",},
			{"1992-12-19 11:59:11", "1992-12-19 11:59:59",},
		}
		for _,v := range testCase {
			as.Equal(gtime.LastMinute(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}


func TestFirstHour(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-19 00:01:11", "1992-12-19 00:00:00",},
			{"1992-12-19 01:01:11", "1992-12-19 00:00:00",},
			{"1992-12-19 02:02:11", "1992-12-19 00:00:00",},
			{"1992-12-19 22:58:11", "1992-12-19 00:00:00",},
			{"1992-12-19 23:59:11", "1992-12-19 00:00:00",},
		}
		for _,v := range testCase {
			as.Equal(gtime.FirstHour(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}

func TestLastHour(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-19 00:01:11", "1992-12-19 23:59:59",},
			{"1992-12-19 01:01:11", "1992-12-19 23:59:59",},
			{"1992-12-19 02:02:11", "1992-12-19 23:59:59",},
			{"1992-12-19 22:58:11", "1992-12-19 23:59:59",},
			{"1992-12-19 23:59:11", "1992-12-19 23:59:59",},
		}
		for _,v := range testCase {
			as.Equal(gtime.LastHour(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}

func TestFirstDay(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-12-01 00:01:11", "1992-12-01 00:00:00",},
			{"1992-12-02 01:01:11", "1992-12-01 00:00:00",},
			{"1992-12-03 02:02:11", "1992-12-01 00:00:00",},
			{"1992-12-30 22:58:11", "1992-12-01 00:00:00",},
			{"1992-12-31 23:59:11", "1992-12-01 00:00:00",},
		}
		for _,v := range testCase {
			as.Equal(gtime.FirstDay(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}

func TestLastDay(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-01-19 00:01:11", "1992-01-31 23:59:59",},
			{"1992-03-19 01:01:11", "1992-03-31 23:59:59",},
			{"1992-04-19 22:58:11", "1992-04-30 23:59:59",},
			{"1992-05-19 23:59:11", "1992-05-31 23:59:59",},
			{"1992-06-19 23:59:11", "1992-06-30 23:59:59",},
			{"1992-07-19 23:59:11", "1992-07-31 23:59:59",},
			{"1992-08-19 23:59:11", "1992-08-31 23:59:59",},
			{"1992-09-19 23:59:11", "1992-09-30 23:59:59",},
			{"1992-10-19 23:59:11", "1992-10-31 23:59:59",},
			{"1992-11-19 23:59:11", "1992-11-30 23:59:59",},
			{"1992-12-19 23:59:11", "1992-12-31 23:59:59",},

			// 各种闰月
			{"1952-02-19 02:02:11", "1952-02-29 23:59:59",},
			{"1956-02-19 02:02:11", "1956-02-29 23:59:59",},
			{"1960-02-19 02:02:11", "1960-02-29 23:59:59",},
			{"1964-02-19 02:02:11", "1964-02-29 23:59:59",},
			// 非闰月
			{"1965-02-19 02:02:11", "1965-02-28 23:59:59",},
			{"1967-02-19 02:02:11", "1967-02-28 23:59:59",},
		}
		for _,v := range testCase {
			as.Equal(gtime.LastDay(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}

func TestFirstMonth(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-01-01 00:01:11", "1992-01-01 00:00:00",},
			{"1992-02-02 01:01:11", "1992-01-01 00:00:00",},
			{"1992-11-03 02:02:11", "1992-01-01 00:00:00",},
			{"1992-12-30 22:58:11", "1992-01-01 00:00:00",},
		}
		for _,v := range testCase {
			as.Equal(gtime.FirstMonth(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}

func TestLastMonth(t *testing.T) {
	as := gtest.NewAS(t)
	{
		testCase := [][]string{
			{"1992-01-19 00:01:11", "1992-12-31 23:59:59",},
			{"1992-03-19 01:01:11", "1992-12-31 23:59:59",},
			{"1992-04-19 22:58:11", "1992-12-31 23:59:59",},
			{"1992-05-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-06-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-07-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-08-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-09-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-10-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-11-19 23:59:11", "1992-12-31 23:59:59",},
			{"1992-12-19 23:59:11", "1992-12-31 23:59:59",},
		}
		for _,v := range testCase {
			as.Equal(gtime.LastMonth(gtime.ParseUTC(gtime.LayoutTime, v[0])).Format(gtime.LayoutTime), v[1])
		}
	}
}