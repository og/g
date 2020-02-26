package gtime

import (
	"time"
)

type Range struct {
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