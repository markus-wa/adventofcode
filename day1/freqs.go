package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	ints := parse(string(input))
	fmt.Println("Result:", calibrate(ints))
	fmt.Println("Twice:", twice(ints))
}

func parse(input string) []int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var res []int
	for scanner.Scan() {
		f, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		res = append(res, f)
	}
	return res
}

func calibrate(input []int) int {
	var res int
	for _, f := range input {
		res += f
	}

	return res
}

func twice(input []int) int {
	freqs := make(map[int]struct{})
	var res int
	for {
		for _, f := range input {
			res += f
			if _, ok := freqs[res]; ok {
				return res
			}
			freqs[res] = struct{}{}
		}
	}
}
