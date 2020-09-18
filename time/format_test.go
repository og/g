package gtime

import (
	gtest "github.com/og/x/test"
	"testing"
	"time"
)

func TestFormatChinaTime(t *testing.T) {
	as := gtest.NewAS(t)
	sometime, err := time.Parse(time.RFC3339, "1992-12-31T23:00:00+00:00")
	as.NoError(err)
	{
		as.Equal(FormatChinaTime(sometime), "1993-01-01 07:00:00")
	}
	{
		as.Equal(FormatChinaDate(sometime), "1993-01-01")
	}
	{
		as.Equal(FormatChinaYear(sometime), "1993")
	}
	{
		as.Equal(FormatChinaYearAndMonth(sometime), "1993-01")
	}
}
