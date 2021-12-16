package timefactoring

import (
	"fmt"
)

type Unit string

const (
	SECONDS Unit = "seconds"
	MINUTES Unit = "minutes"
	HOURS   Unit = "hours"
	DAYS    Unit = "days"
	WEEKS   Unit = "weeks"
)

func MinutesToSeconds(minutes int) int {
	return minutes * 60
}

func HoursToSeconds(hours int) int {
	return MinutesToSeconds(hours * 60)
}

func DaysToSeconds(days int) int {
	return HoursToSeconds(days * 24)
}

func WeeksToSeconds(weeks int) int {
	return DaysToSeconds(weeks * 7)
}

func ToSeconds(number int, unit Unit) (seconds int, err error) {
	switch unit {
	case SECONDS:
		seconds = number
	case MINUTES:
		seconds = MinutesToSeconds(number)
	case HOURS:
		seconds = HoursToSeconds(number)
	case DAYS:
		seconds = DaysToSeconds(number)
	case WEEKS:
		seconds = WeeksToSeconds(number)
	default:
		err = fmt.Errorf("Unit %s is not supported", unit)
	}
	return
}
