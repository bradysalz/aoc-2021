package main

func dayNinePartOne() int {
	data := readFileToIntArray("input/day_09.txt")
	lowPoints := findLowPoints(data)
	partOne := 0
	for _, point := range lowPoints {
		partOne += point
	}
	partOne += len(lowPoints)

	return partOne
}

// Returns a list of cells that are their local minima in X+Y
func findLowPoints(data [][]int) []int {
	// assume rectangular
	sizeX := len(data)
	sizeY := len(data[0])

	lowPoints := make([]int, 0)
	for idx, row := range data {
		for idy, cell := range row {
			if (idx-1) >= 0 && (cell >= data[idx-1][idy]) {
				continue
			} else if (idx+1) < sizeX && (cell >= data[idx+1][idy]) {
				continue
			} else if (idy-1) >= 0 && (cell >= data[idx][idy-1]) {
				continue
			} else if (idy+1) < sizeY && (cell >= data[idx][idy+1]) {
				continue
			} else {
				lowPoints = append(lowPoints, cell)
			}
		}
	}
	return lowPoints
}
