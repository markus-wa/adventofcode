package day3

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"
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
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	actual := parse(input)

	expected := []r2.Rect{
		{X: r1.Interval{Lo: 1 + 1, Hi: 2 + 4}, Y: r1.Interval{Lo: 3 + 1, Hi: 4 + 4}},
		{X: r1.Interval{Lo: 3 + 1, Hi: 4 + 4}, Y: r1.Interval{Lo: 1 + 1, Hi: 2 + 4}},
		{X: r1.Interval{Lo: 5 + 1, Hi: 6 + 2}, Y: r1.Interval{Lo: 5 + 1, Hi: 6 + 2}},
	}
	assert.Equal(t, expected, actual)
}

func TestOverlapping(t *testing.T) {
	input := []r2.Rect{
		{X: r1.Interval{Lo: 1 + 1, Hi: 2 + 4 + 1}, Y: r1.Interval{Lo: 3 + 1, Hi: 4 + 4}},
		{X: r1.Interval{Lo: 3 + 1, Hi: 4 + 4}, Y: r1.Interval{Lo: 1 + 1, Hi: 2 + 4 + 1}},
		{X: r1.Interval{Lo: 5 + 1, Hi: 6 + 2}, Y: r1.Interval{Lo: 5 + 1, Hi: 6 + 2}},
	}

	actual := overlapping(input)

	expected := 11
	assert.Equal(t, expected, actual)
}

func TestOverlappingInput(t *testing.T) {
	strs := parse(inputTxt)

	actual := overlapping(strs)

	expected := 110195
	assert.Equal(t, expected, actual)
}

func TestNonOverlapping(t *testing.T) {
	input := []r2.Rect{
		{X: r1.Interval{Lo: 1 + 1, Hi: 2 + 4}, Y: r1.Interval{Lo: 3 + 1, Hi: 4 + 4}},
		{X: r1.Interval{Lo: 3 + 1, Hi: 4 + 4}, Y: r1.Interval{Lo: 1 + 1, Hi: 2 + 4}},
		{X: r1.Interval{Lo: 5 + 1, Hi: 6 + 2}, Y: r1.Interval{Lo: 5 + 1, Hi: 6 + 2}},
	}

	actual := nonOverlapping(input)

	expected := 3
	assert.Equal(t, expected, actual)
}

func TestNonOverlappingInput(t *testing.T) {
	strs := parse(inputTxt)

	actual := nonOverlapping(strs)

	expected := 894
	assert.Equal(t, expected, actual)
}
