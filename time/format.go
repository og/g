package gtime

import "time"

func FormatChinaTime(t time.Time) string {
	return t.In(LocChina).Format(LayoutTime)
}
func FormatChinaDate(t time.Time) string {
	return t.In(LocChina).Format(LayoutDate)
}
func FormatChinaYear(t time.Time) string {
	return t.In(LocChina).Format(LayoutYear)
}
func FormatChinaYearAndMonth(t time.Time) string {
	return t.In(LocChina).Format(LayoutYearAndMonth)
}