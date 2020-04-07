package gtime_test

import (
	gtime "github.com/og/x/time"
	"testing"
)

func TestParse(t *testing.T) {
	{
		// 本机时间要是北京时间
		v := gtime.ParseChina(gtime.Second, "2020-04-08 00:00:00");
		if v.Hour() != 0 {
			panic("china 1 error hour")
		}
		if v.UTC().Hour() != 16 {
			panic("china 2 error hour")
		}
	}
	{
		// 本机时间要是北京时间
		v := gtime.ParseUTC(gtime.Second, "2020-04-08 00:00:00");
		if v.Hour() != 0 {
			panic("utc 1 error hour")
		}
		if v.Local().Hour() != 8 {
			panic("utc 2 error hour")
		}
	}
}
