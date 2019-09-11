package gtime

import (
	"time"
)

func GetDayRange (timeRange Range) (days []string) {
	days = []string{}
	var beyondEndTime time.Time
	switch timeRange.Type {
	case Dict().Range.Type.Year:
		beyondEndTime = timeRange.End.AddDate(1,0,0)
	case Dict().Range.Type.Month:
		beyondEndTime = timeRange.End.AddDate(0,1,0)
	case Dict().Range.Type.Day:
		beyondEndTime = timeRange.End.AddDate(0,0,1)
	}

	var dayTimeList []time.Time
	nextDay := timeRange.Start
	for ; nextDay.Before(beyondEndTime) ; nextDay = nextDay.AddDate(0,0,1) {
		dayTimeList = append(dayTimeList, nextDay)
	}
	for _, dayTime := range dayTimeList {
		days = append(days, dayTime.Format(Day))
	}
	return
}
