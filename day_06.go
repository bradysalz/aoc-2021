package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func daySixPartOne() int {
	const NUM_DAYS = 80
	state := readInputDaySix("input/day_06.txt")

	for d := 1; d < NUM_DAYS; d++ {
		for s := range state {
			state[s] -= 1
			if state[s] == 0 {
				state = append(state, 9)
				state[s] = 7
			}
		}
	}
	return len(state)
}

func daySixPartTwo() int {
	// full disclosure yes i got wrecked thinking it was just one constant to update in part two
	const NUM_DAYS = 256
	const MAX_LIFETIME = 9
	total := 0
	state := readInputDaySix("input/day_06.txt")
	lifetimes := make([]int, MAX_LIFETIME+1)
	for _, fish := range state {
		lifetimes[fish] += 1
	}
	for d := 1; d <= NUM_DAYS; d++ {
		lifetimes[7] += lifetimes[0]
		lifetimes[9] += lifetimes[0]
		for l := 0; l < MAX_LIFETIME; l++ {
			lifetimes[l] = lifetimes[l+1]
		}
		lifetimes[9] = 0
		total = 0
		for l := range lifetimes {
			total += lifetimes[l]
		}
	}

	return total
}

func readInputDaySix(path string) []int {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	retVal := make([]int, 0)
	for _, line := range strings.Split(string(lines), "\n") {
		for _, word := range strings.Split(line, ",") {
			val, err := strconv.ParseInt(word, 0, 0)
			if err != nil {
				panic(err)
			}
			retVal = append(retVal, int(val))
		}
	}
	return retVal
}
