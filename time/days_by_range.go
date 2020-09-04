package gtime

import (
	"time"
)

func DaysByRange (timeRange Range, daysLoc *time.Location) (days []string) {
	days = []string{}
	var beyondEndTime time.Time
	timeRange.Type.Switch(
		func(_year int) {
			beyondEndTime = timeRange.End.AddDate(1,0,0)
		},
		func(_month bool) {
			beyondEndTime = timeRange.End.AddDate(0,1,0)
		},
		func(_day string) {
			beyondEndTime = timeRange.End.AddDate(0,0,1)
		},
	)
	var dayTimeList []time.Time
	nextDay := timeRange.Start
	for ; nextDay.Before(beyondEndTime) ; nextDay = nextDay.AddDate(0,0,1) {
		dayTimeList = append(dayTimeList, nextDay)
	}
	for _, dayTime := range dayTimeList {
		days = append(days, dayTime.In(daysLoc).Format(LayoutDate))
	}
	return
}
