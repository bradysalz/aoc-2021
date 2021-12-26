package main

const ARR_SIZE = 10

// TIL lower-case == private, upper-case == public
type cell struct {
	value   int
	flashed bool
}

func dayElevenPartOne() int {
	data := readFileToIntArray("input/day_11.txt")
	array := convertToCellArray(data)

	totalFlashes := 0
	for step := 0; step < 100; step++ {
		incrementArray(&array)
		lastRound := calculateFlashes(&array, 0)
		resetFlashes(&array)
		totalFlashes += lastRound
	}
	return totalFlashes
}

func dayEleventPartTwo() int {
	data := readFileToIntArray("input/day_11.txt")
	array := convertToCellArray(data)

	day := 1
	for {
		incrementArray(&array)
		lastRound := calculateFlashes(&array, 0)
		resetFlashes(&array)
		if lastRound == (ARR_SIZE * ARR_SIZE) {
			return day
		}
		day++
	}
}

func calculateFlashes(data *[][]cell, currFlashSum int) int {
	flashes := currFlashSum
	for idx, row := range *data {
		for idy, square := range row {
			// can only flash once per round
			if (*data)[idx][idy].flashed {
				continue
			}
			if square.value > 9 {
				updateAdjacentCells(data, idx, idy)
				(*data)[idx][idy].flashed = true
				(*data)[idx][idy].value = 0
				flashes = calculateFlashes(data, 1+currFlashSum)
			}
		}
	}
	return flashes
}

func resetFlashes(data *[][]cell) {
	for idx, row := range *data {
		for idy, square := range row {
			if square.flashed {
				(*data)[idx][idy].value = 0
				(*data)[idx][idy].flashed = false
			}
		}
	}
}

func convertToCellArray(data [][]int) [][]cell {
	arr := make([][]cell, ARR_SIZE)
	for a := range arr {
		arr[a] = make([]cell, ARR_SIZE)
	}

	for x := range data {
		for y := range data[x] {
			arr[x][y] = cell{value: data[x][y], flashed: false}
		}
	}
	return arr
}

func updateAdjacentCells(data *[][]cell, x int, y int) {
	neighborCoords := [][]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
		{x - 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y - 1},
		{x + 1, y + 1},
	}
	for _, coord := range neighborCoords {
		if (coord[0] < 0) || (coord[0] >= ARR_SIZE) || (coord[1] < 0) || (coord[1] >= ARR_SIZE) {
			continue
		}
		(*data)[coord[0]][coord[1]].value += 1
	}
}

func incrementArray(data *[][]cell) {
	for idx, row := range *data {
		for idy := range row {
			(*data)[idx][idy].value += 1
		}
	}
}
