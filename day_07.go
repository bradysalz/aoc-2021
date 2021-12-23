package main

import "math"

func daySevenPartOne(isPartTwo bool) int {
	// yay re-usability
	state := readInputDaySix("input/day_07.txt")
	maxValue, _ := getMaxValueAndIndex(state)

	cost := 0
	costArr := make([]int, 0)
	for v := 0; v <= maxValue; v++ {
		cost = 0
		for _, s := range state {
			if isPartTwo {
				for idx := 0; idx <= int(math.Abs(float64(s-v))); idx++ {
					cost -= idx
				}
			} else {
				cost -= int(math.Abs(float64(s - v)))
			}
		}
		costArr = append(costArr, cost)
	}
	bestFuel, _ := getMaxValueAndIndex(costArr)
	return bestFuel * -1
}
