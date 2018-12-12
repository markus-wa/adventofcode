package day4

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputTxt string

func TestMain(m *testing.M) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	inputTxt = string(input)

	os.Exit(m.Run())
}

func TestParse(t *testing.T) {
	input := `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

	actual := shiftsFromEvents(parse(input))

	expected := []shift{
		{guardID: 10, naps: []nap{{start: 5, end: 25}, {start: 30, end: 55}}},
		{guardID: 99, naps: []nap{{start: 40, end: 50}}},
		{guardID: 10, naps: []nap{{start: 24, end: 29}}},
		{guardID: 99, naps: []nap{{start: 36, end: 46}}},
		{guardID: 99, naps: []nap{{start: 45, end: 55}}},
	}
	assert.Equal(t, expected, actual)
}

func TestSleepiestGuard(t *testing.T) {
	input := []shift{
		{guardID: 10, naps: []nap{{start: 5, end: 25}, {start: 30, end: 55}}},
		{guardID: 99, naps: []nap{{start: 40, end: 50}}},
		{guardID: 10, naps: []nap{{start: 24, end: 29}}},
		{guardID: 99, naps: []nap{{start: 36, end: 46}}},
		{guardID: 99, naps: []nap{{start: 45, end: 55}}},
	}

	actual := sleepiestGuardTimesMins(input)

	expected := 10 * 24
	assert.Equal(t, expected, actual)
}

func TestSleepiestGuardInput(t *testing.T) {
	shifts := shiftsFromEvents(parse(inputTxt))

	actual := sleepiestGuardTimesMins(shifts)

	expected := 99759
	assert.Equal(t, expected, actual)
}
func TestSameMinSleepGuardTimesMin(t *testing.T) {
	input := []shift{
		{guardID: 10, naps: []nap{{start: 5, end: 25}, {start: 30, end: 55}}},
		{guardID: 99, naps: []nap{{start: 40, end: 50}}},
		{guardID: 10, naps: []nap{{start: 24, end: 29}}},
		{guardID: 99, naps: []nap{{start: 36, end: 46}}},
		{guardID: 99, naps: []nap{{start: 45, end: 55}}},
	}

	actual := sameMinSleepGuardTimesMin(input)

	expected := 99 * 45
	assert.Equal(t, expected, actual)
}

func TestSameMinSleepGuardTimesMinInput(t *testing.T) {
	shifts := shiftsFromEvents(parse(inputTxt))

	actual := sameMinSleepGuardTimesMin(shifts)

	expected := 97884
	assert.Equal(t, expected, actual)
}
