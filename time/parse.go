package gtime

import (
	"time"
)

// use time.Parse, when parse error panic(err)
func Parse(layout string, value string) time.Time {
	v, err := time.Parse(layout, value)
	if err != nil { panic(err) }
	return v
}
