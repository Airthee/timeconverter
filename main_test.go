package timefactoring_test

import (
	"testing"

	timefactoring "github.com/airthee/time-converter"
	"github.com/stretchr/testify/assert"
)

func TestSecondesToSeconds(t *testing.T) {
	// ARRANGE
	input := 45
	expected := 45

	// ACT
	result, _ := timefactoring.ToSeconds(input, "seconds")

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestMinutesToSeconds(t *testing.T) {
	// ARRANGE
	input := 45
	expected := 45 * 60

	// ACT
	result, _ := timefactoring.ToSeconds(input, "minutes")

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestHoursToSeconds(t *testing.T) {
	// ARRANGE
	input := 35
	expected := 35 * 3600

	// ACT
	result, _ := timefactoring.ToSeconds(input, "hours")

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestDaysToSeconds(t *testing.T) {
	// ARRANGE
	input := 35
	expected := 35 * 3600 * 24

	// ACT
	result, _ := timefactoring.ToSeconds(input, "days")

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestWeeksToSeconds(t *testing.T) {
	// ARRANGE
	input := 35
	expected := 35 * 3600 * 24 * 7

	// ACT
	result, _ := timefactoring.ToSeconds(input, "weeks")

	// ASSERT
	assert.Equal(t, expected, result)
}

func TestShouldReturnErrorIfWrongUnit(t *testing.T) {
	// ACT
	_, err := timefactoring.ToSeconds(15, "holidays")

	// ASSERT
	assert.EqualError(t, err, "Unit holidays is not supported")
}
