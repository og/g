package gtime

import (
	"time"
)

type rangeType struct {
	value string
}
type Range struct {
	// 注意  Type Start End 都不要变，因为 og/gofree 的 f.TimeRange 依赖了 Range的这三个字段
	Type rangeType
	Start time.Time
	End time.Time
}
func (self Range) Dict() (dict struct{
	Type struct{
		Year rangeType
		Month rangeType
		Day rangeType
	}
}) {
	dict.Type.Year = rangeType{"year"}
	dict.Type.Month = rangeType{"month"}
	dict.Type.Day = rangeType{"day"}
	return
}