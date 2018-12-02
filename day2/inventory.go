package day2

import (
	"bufio"
	"strings"
)

func parse(input string) []string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func checksum(input []string) int {
	var nTwice, nTrice int
	for _, s := range input {
		occurences := make(map[rune]int)
		for _, c := range s {
			occurences[c] = occurences[c] + 1
		}
		var twice, trice bool
		for _, v := range occurences {
			twice = twice || v == 2
			trice = trice || v == 3
		}
		if twice {
			nTwice++
		}
		if trice {
			nTrice++
		}
	}
	return nTwice * nTrice
}

func similarIDs(input []string) string {
	shortIDs := make(map[string]string)
	for i := 0; ; i++ {
		for _, str := range input {
			short := str[:i] + str[i+1:]
			if full, ok := shortIDs[short]; ok && full != str {
				return short
			}
			shortIDs[short] = str
		}
	}
}
