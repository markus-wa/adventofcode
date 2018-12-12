package day4

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
	"time"
)

type shift struct {
	guardID int
	naps    []nap
}

type nap struct {
	start int
	end   int
}

func parse(input string) (events []event) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		events = append(events, parseEvent(scanner.Text()))
	}

	return
}

type events []event

func (s events) Len() int {
	return len(s)
}
func (s events) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type byTime struct{ events }

func (s byTime) Less(i, j int) bool {
	return s.events[i].time.Unix() < s.events[j].time.Unix()
}

func shiftsFromEvents(events []event) (shifts []shift) {
	sort.Sort(byTime{events})

	var currentShift shift
	var currentNap nap
	for _, e := range events {
		switch e.action {
		case "falls asleep":
			currentNap.start = e.time.Minute()
		case "wakes up":
			currentNap.end = e.time.Minute()
			currentShift.naps = append(currentShift.naps, currentNap)
		default:
			if currentShift.guardID != 0 {
				shifts = append(shifts, currentShift)
			}

			guardID, err := strconv.Atoi(strings.TrimRight(strings.Split(strings.Split(e.action, "#")[1], "begins shift")[0], " "))
			checkError(err)
			currentShift = shift{guardID: guardID}
		}
	}
	shifts = append(shifts, currentShift)
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type event struct {
	time   time.Time
	action string
}

func parseEvent(text string) (ev event) {
	text = strings.TrimLeft(text, "[")
	spl := strings.Split(text, "] ")
	timeStr := spl[0]
	action := spl[1]
	time, err := time.Parse("2006-01-02 15:04", timeStr)
	if err != nil {
		panic(err)
	}
	return event{
		time:   time,
		action: action,
	}
}

type guard struct {
	id             int
	totalSleepMins int
	naps           []nap
}

func sleepiestGuardTimesMins(shifts []shift) int {
	guardIDtoGuard := make(map[int]*guard)

	for _, s := range shifts {
		if _, ok := guardIDtoGuard[s.guardID]; !ok {
			guardIDtoGuard[s.guardID] = &guard{id: s.guardID}
		}
		currentGuard := guardIDtoGuard[s.guardID]
		currentGuard.naps = append(guardIDtoGuard[s.guardID].naps, s.naps...)
		for _, n := range s.naps {
			currentGuard.totalSleepMins = currentGuard.totalSleepMins + n.end - n.start
		}
	}
	var maxMinsGuard guard
	for _, g := range guardIDtoGuard {
		if maxMinsGuard.totalSleepMins < g.totalSleepMins {
			maxMinsGuard = *g
		}
	}
	minsToTimesAsleep := make(map[int]int)
	for _, n := range maxMinsGuard.naps {
		for i := n.start; i < n.end; i++ {
			minsToTimesAsleep[i] = minsToTimesAsleep[i] + 1
		}
	}
	var maxMin, maxTimesAsleep int
	for min, timesAsleep := range minsToTimesAsleep {
		if maxTimesAsleep < timesAsleep {
			maxTimesAsleep = timesAsleep
			maxMin = min
		}
	}
	return maxMinsGuard.id * maxMin
}

func sameMinSleepGuardTimesMin(shifts []shift) int {
	guardIDtoMinsToTimesAsleep := make(map[int]map[int]int)

	for _, s := range shifts {
		if _, ok := guardIDtoMinsToTimesAsleep[s.guardID]; !ok {
			guardIDtoMinsToTimesAsleep[s.guardID] = make(map[int]int)
		}

		for _, n := range s.naps {
			for min := n.start; min < n.end; min++ {
				guardIDtoMinsToTimesAsleep[s.guardID][min] = guardIDtoMinsToTimesAsleep[s.guardID][min] + 1
			}
		}
	}

	var guardID, maxAsleepMin, maxTimesAsleep int
	for gid, minsToTimesAsleep := range guardIDtoMinsToTimesAsleep {
		for min, timesAsleep := range minsToTimesAsleep {
			if maxTimesAsleep < timesAsleep {
				maxTimesAsleep = timesAsleep
				maxAsleepMin = min
				guardID = gid
			}
		}
	}
	return guardID * maxAsleepMin
}
