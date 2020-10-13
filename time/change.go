package gtime

import (
	"time"
)
// 1992-12-19 11:11:11.988 => 1992-12-19 11:11:00.000
func FirstNanoseconds (v time.Time) time.Time{
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute() ,v.Second(), 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-19 11:11:00
func FirstSecond (v time.Time) time.Time{
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute() ,0, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-19 11:11:59
func LastSecond (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute() ,59, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-19 11:00:00
func FirstMinute (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), 0 ,0, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-19 11:59:59
func LastMinute (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), 59 ,59, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-19 00:00:00
func FirstHour (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), 0, 0 ,0, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-19 23:59:59
func LastHour (v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), 23, 59 ,59, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-01 00:00:00
func FirstDay(v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), 1, 0, 0 ,0, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-31 23:59:59
func LastDay(v time.Time) time.Time {
	return time.Date(v.Year(), v.Month()+1, 1, 23, 59 ,59, 0, v.Location()).AddDate(0,0,-1)
}
// 1992-12-19 11:11:11 => 1992-01-01 00:00:00
func FirstMonth(v time.Time) time.Time {
	return time.Date(v.Year(),1, 1, 0, 0 ,0, 0, v.Location())
}
// 1992-12-19 11:11:11 => 1992-12-31 23:59:59
func LastMonth(v time.Time) time.Time {
	return time.Date(v.Year(), time.December, 31, 23, 59 ,59, 0, v.Location())
}