package gtime

import (
	"time"
)

// use time.ParseInLocation(layout value, time.UTC), when parse error panic(err)
func ParseUTC(layout string, value string) time.Time {
	v, err := time.ParseInLocation(layout, value, time.UTC)
	if err != nil { panic(err) }
	return v
}

// 除非执行一些定时任务,或者 value 是写死的,否则不要使用 ParseChina 而是 Parse UTC
func ParseChina(layout string, value string) time.Time {
	v, err := time.ParseInLocation(layout, value, LocChina)
	if err != nil { panic(err) }
	return v
}
func NowChina() time.Time {
	return time.Now().In(LocChina)
}
func NowUTC() time.Time {
	return time.Now().In(time.UTC)
}