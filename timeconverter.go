package timeconverter

type FactorizedTime struct {
	Weeks   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

const (
	SECONDS_PER_MINUTES = 60
	SECONDS_PER_HOURS   = SECONDS_PER_MINUTES * 60
	SECONDS_PER_DAYS    = SECONDS_PER_HOURS * 24
	SECONDS_PER_WEEKS   = SECONDS_PER_DAYS * 7
)

func MinutesToSeconds(minutes int) int {
	return minutes * SECONDS_PER_MINUTES
}

func HoursToSeconds(hours int) int {
	return hours * SECONDS_PER_HOURS
}

func DaysToSeconds(days int) int {
	return days * SECONDS_PER_DAYS
}

func WeeksToSeconds(weeks int) int {
	return weeks * SECONDS_PER_WEEKS
}

func ToSeconds(weeks int, days int, hours int, minutes int, seconds int) int {
	return WeeksToSeconds(weeks) + DaysToSeconds(days) + HoursToSeconds(hours) + MinutesToSeconds(minutes) + seconds
}

func SecondsToMinutes(seconds int) float64 {
	return float64(seconds) / SECONDS_PER_MINUTES
}

func SecondsToHours(seconds int) float64 {
	return float64(seconds) / SECONDS_PER_HOURS
}

func SecondsToDays(seconds int) float64 {
	return float64(seconds) / SECONDS_PER_DAYS
}

func SecondsToWeeks(seconds int) float64 {
	return float64(seconds) / SECONDS_PER_WEEKS
}

func Factorize(seconds int) FactorizedTime {
	weeks := int(SecondsToWeeks(seconds))
	seconds -= weeks * SECONDS_PER_WEEKS
	days := int(SecondsToDays(seconds))
	seconds -= days * SECONDS_PER_DAYS
	hours := int(SecondsToHours(seconds))
	seconds -= hours * SECONDS_PER_HOURS
	minutes := int(SecondsToMinutes(seconds))
	seconds -= minutes * SECONDS_PER_MINUTES

	return FactorizedTime{weeks, days, hours, minutes, seconds}
}
