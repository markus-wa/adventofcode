package main

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
	input := `+1
-2
+5`

	actual := parse(input)

	expected := []int{1, -2, 5}
	assert.Equal(t, expected, actual)
}

func TestCalibrate(t *testing.T) {
	input := []int{1, -2, 5}

	actual := calibrate(input)

	expected := 4
	assert.Equal(t, expected, actual)
}
func TestTwice(t *testing.T) {
	input := []int{1, -2, 5, 2, -7}

	actual := twice(input)

	expected := -1
	assert.Equal(t, expected, actual)
}

func TestParseAndCalibrate(t *testing.T) {
	ints := parse(inputTxt)

	actual := calibrate(ints)

	expected := 479
	assert.Equal(t, expected, actual)
}

func TestParseAndTwice(t *testing.T) {
	ints := parse(inputTxt)

	actual := twice(ints)

	expected := 66105
	assert.Equal(t, expected, actual)
}

func TestMainFunc(t *testing.T) {
	main()
}
