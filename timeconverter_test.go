package timeconverter_test

import (
	"testing"

	"github.com/airthee/timeconverter"
	"github.com/stretchr/testify/assert"
)

func TestMinutesToSeconds(t *testing.T) {
	// ARRANGE
	input := 45
	expected := 45 * 60

	// ACT
	result := timeconverter.MinutesToSeconds(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestHoursToSeconds(t *testing.T) {
	// ARRANGE
	input := 35
	expected := 35 * 3600

	// ACT
	result := timeconverter.HoursToSeconds(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestDaysToSeconds(t *testing.T) {
	// ARRANGE
	input := 35
	expected := 35 * 3600 * 24

	// ACT
	result := timeconverter.DaysToSeconds(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestWeeksToSeconds(t *testing.T) {
	// ARRANGE
	input := 35
	expected := 35 * 3600 * 24 * 7

	// ACT
	result := timeconverter.WeeksToSeconds(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestSecondsToMinutes(t *testing.T) {
	// ARRANGE
	input := 90
	expected := 1.5

	// ACT
	result := timeconverter.SecondsToMinutes(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestSecondsToHours(t *testing.T) {
	// ARRANGE
	input := (3600 * 2) + (60 * 45)
	expected := 2.75

	// ACT
	result := timeconverter.SecondsToHours(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestSecondsToDays(t *testing.T) {
	// ARRANGE
	input := (3600 * 24 * 3) + (3600 * 6)
	expected := 3.25

	// ACT
	result := timeconverter.SecondsToDays(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestSecondsToWeeks(t *testing.T) {
	// ARRANGE
	input := 3600 * 24 * 7
	expected := 1.

	// ACT
	result := timeconverter.SecondsToWeeks(input)

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestSecondsToFactorized(t *testing.T) {
	assert := assert.New(t)

	// ARRANGE
	numberOfWeeks := 5
	numberOfDays := 3
	numberOfHours := 23
	numberOfMinutes := 41
	numberOfSeconds := 34
	totalNumberOfSeconds := timeconverter.ToSeconds(numberOfWeeks, numberOfDays, numberOfHours, numberOfMinutes, numberOfSeconds)

	// ACT
	result := timeconverter.Factorize(totalNumberOfSeconds)

	// ASSERT
	assert.Equal(numberOfWeeks, result.Weeks)
	assert.Equal(numberOfDays, result.Days)
	assert.Equal(numberOfHours, result.Hours)
	assert.Equal(numberOfMinutes, result.Minutes)
	assert.Equal(numberOfSeconds, result.Seconds)
}
