package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	LastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()
	diffWeek := wDay - firstDay.Weekday()
	if diffWeek < 0 {
		diffWeek += 7
	}
	switch wSched {
	case First:
		return int(diffWeek) + 1
	case Second:
		return int(diffWeek) + 8
	case Third:
		return int(diffWeek) + 15
	case Fourth:
		return int(diffWeek) + 22
	case Last:
		if LastDay >= int(diffWeek)+29 {
			return int(diffWeek) + 29
		}
		return int(diffWeek) + 22
	case Teenth:
		TeenthDay := time.Date(year, month, 13, 0, 0, 0, 0, time.Local).Weekday()
		diffDay := wDay - TeenthDay
		if diffDay < 0 {
			diffDay += 7
		}
		return int(diffDay) + 13
	default:
		return 0
	}
}
