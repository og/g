package gtime

import (
	"errors"
	"time"
)

type RangeType string
func (v RangeType) String() string {
	return string(v)
}
type Range struct {
	// 注意  Type Start End 都不要变，因为 og/gofree 的 f.TimeRange 依赖了 Range的这三个字段
	Type RangeType
	Start time.Time
	End time.Time
}
func (RangeType) Enum() (enum struct{
	Year RangeType
	Month RangeType
	Day RangeType
}) {
	enum.Year = RangeType("year")
	enum.Month = RangeType("month")
	enum.Day = RangeType("day")
	return
}
func (v RangeType) Switch(
	Year func(_year int),
	Month func(_month bool),
	Day func(_day string),
	) {
	enum := v.Enum()
	switch v {
	default:
		s := v.String()
		if len(s) == 0 {
			s = "empty string"
		}
		panic(errors.New("RangeType: can not be (" + s + ")"))
	case enum.Year:
		Year(0)
	case enum.Month:
		Month(false)
	case enum.Day:
		Day("")
	}
}