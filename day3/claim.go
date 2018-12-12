package day3

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/geo/r2"
)

func parse(input string) (res []r2.Rect) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		res = append(res, parseRect(scanner.Text()))
	}

	return
}

func parseRect(text string) (rect r2.Rect) {
	text = strings.Trim(strings.Split(text, "@")[1], " ")
	spl := strings.Split(text, ":")
	xy := strings.Split(strings.Trim(spl[0], " "), ",")
	wh := strings.Split(strings.Trim(spl[1], " "), "x")

	var err error
	rect.X.Lo, err = strconv.ParseFloat(xy[0], 64)
	checkErr(err)
	rect.Y.Lo, err = strconv.ParseFloat(xy[1], 64)
	checkErr(err)
	rect.X.Hi, err = strconv.ParseFloat(wh[0], 64)
	checkErr(err)
	rect.Y.Hi, err = strconv.ParseFloat(wh[1], 64)
	checkErr(err)
	// offset != start of rect
	rect.X.Lo++
	rect.Y.Lo++
	// we have x,y WxH but we need Lo/Hi
	rect.X.Hi += rect.X.Lo
	rect.Y.Hi += rect.Y.Lo

	return
}

func overlapping(rects []r2.Rect) (nOverlapping int) {
	overlapping := make(map[int]map[int]bool)
	for i, r := range rects {
		for _, r2 := range rects[i+1:] {
			if r.Intersects(r2) {
				inters := r.Intersection(r2)
				for i := int(inters.X.Lo); i < int(inters.X.Hi); i++ {
					if _, ok := overlapping[i]; !ok {
						overlapping[i] = make(map[int]bool)
					}
					for j := int(inters.Y.Lo); j < int(inters.Y.Hi); j++ {
						overlapping[i][j] = true
					}
				}
			}
		}
	}

	for _, v := range overlapping {
		nOverlapping += len(v)
	}
	return
}
func nonOverlapping(rects []r2.Rect) int {
	set := make(map[int]int)
	for i, r := range rects {
		for j, r2 := range rects[i+1:] {
			if r.InteriorIntersects(r2) {
				set[i] = set[i] + 1
				set[i+1+j] = set[i+1+j] + 1
			}
		}
		if set[i] == 0 {
			return i + 1
		}
	}
	fmt.Println(set)

	return -1
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
