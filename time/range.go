package gtime

import (
	"time"
)

type Range struct {
	// 注意  Type Start End 都不要变，因为 og/gofree 的 f.TimeRange 依赖了 Range的这三个字段
	Type string
	Start time.Time
	End time.Time
}
func (self Range) Dict() (dict struct{
	Type struct{
		Year string
		Month string
		Day string
	}
}) {
	dict.Type.Year = "year"
	dict.Type.Month = "month"
	dict.Type.Day = "day"
	return
}