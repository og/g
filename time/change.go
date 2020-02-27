package gtime

import (
	"time"
)


func StartOfSecond (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute() ,0, 0, v.Location())
}
func EndOfSecond (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute() ,59, 0, v.Location())
}
func StartOfMinute (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), 0 ,0, 0, v.Location())
}
func EndOfMinute (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), 59 ,59, 0, v.Location())
}

func StartOfHour (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), 0, 0 ,0, 0, v.Location())
}
func EndOfHour (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), 23, 59 ,59, 0, v.Location())
}
func StartOfDay(v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), 1, 0, 0 ,0, 0, v.Location())
}
func EndOfDay(v time.Time) time.Time {
	return time.Date(v.Year(), v.Month()+1, 1, 23, 59 ,59, 0, v.Location()).AddDate(0,0,-1)
}

func StartOfMonth(v time.Time) time.Time {
	return time.Date(v.Year(),1, 1, 0, 0 ,0, 0, v.Location())
}
func EndOfMonth(v time.Time) time.Time {
	return time.Date(v.Year(), time.December, 31, 23, 59 ,59, 0, v.Location())
}
