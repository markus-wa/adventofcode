package day2

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
	input := `myhposlqgeauywfikztndcvrqr
mbhposlxfeauywoikztndcvjqi
mbhpoulxgeagywfikytndcvjqr
jbhposlxgeauywdikztndcvjqk`

	actual := parse(input)

	expected := []string{"myhposlqgeauywfikztndcvrqr", "mbhposlxfeauywoikztndcvjqi", "mbhpoulxgeagywfikytndcvjqr", "jbhposlxgeauywdikztndcvjqk"}
	assert.Equal(t, expected, actual)
}

func TestChecksum(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}

	actual := checksum(input)

	expected := 12
	assert.Equal(t, expected, actual)
}

func TestChecksumInput(t *testing.T) {
	strs := parse(inputTxt)

	actual := checksum(strs)

	expected := 6474
	assert.Equal(t, expected, actual)
}

func TestSimilarIDs(t *testing.T) {
	input := []string{"aaabb", "abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}

	actual := similarIDs(input)

	expected := "fgij"
	assert.Equal(t, expected, actual)
}

func TestSimilarIDsInput(t *testing.T) {
	strs := parse(inputTxt)

	actual := similarIDs(strs)

	expected := "mxhwoglxgeauywfkztndcvjqr"
	assert.Equal(t, expected, actual)
}
