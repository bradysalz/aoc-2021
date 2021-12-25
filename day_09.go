package main

import (
	"sort"
)

type Cell struct {
	x          int
	y          int
	value      int
	isLowPoint bool
	basin      int
}

func dayNinePartOne() int {
	content := readFileToIntArray("input/day_09.txt")
	lowPoints := findLowPoints(content)
	partOne := 0
	for _, point := range lowPoints {
		partOne += point
	}
	partOne += len(lowPoints)

	return partOne
}

func dayNinePartTwo() int {
	content := readFileToIntArray("input/day_09.txt")
	data := convertIntArrayToCellArray(content)
	findLowPointsOfCells(&data)
	markBasins(&data)
	scores := calculateBasinScores(data)
	sort.Ints(scores)
	return scores[len(scores)-1] * scores[len(scores)-2] * scores[len(scores)-3]
}

func calculateBasinScores(data [][]Cell) []int {
	scoreMap := make(map[int]int)

	for _, row := range data {
		for _, cell := range row {
			_, ok := scoreMap[cell.basin]
			if ok {
				scoreMap[cell.basin] += 1
			} else {
				scoreMap[cell.basin] = 1
			}
		}
	}
	scoreList := make([]int, 0)
	for idx, val := range scoreMap {
		if idx == 0 {
			// ignore "9" cells which have no basin assigned
			continue
		}
		scoreList = append(scoreList, val)
	}
	return scoreList
}

// Iterate through the entire matrix
// If a cell is a LowPoint, BFS out until you hit 9's
func markBasins(data *[][]Cell) {
	basinIdx := 1
	for idx, row := range *data {
		for idy, cell := range row {
			if cell.isLowPoint {
				(*data)[idx][idy].basin = basinIdx
				basinBFS(data, idx, idy, basinIdx)
				basinIdx += 1
			}
		}
	}
}

func basinBFS(data *[][]Cell, startX int, startY int, basinIdx int) {
	// Create the array of visited cells
	sizeX := len(*data)
	sizeY := len((*data)[0])

	visited := make([][]bool, sizeX)
	for idx := 0; idx < sizeX; idx++ {
		visited[idx] = make([]bool, sizeY)
		for idy := 0; idy < sizeY; idy++ {
			visited[idx][idy] = false
		}
	}

	queue := make([]Cell, 0)
	queue = append(queue, (*data)[startX][startY])
	visited[startX][startY] = true

	for len(queue) > 0 {
		// pop off the current cell
		startCell := queue[0]
		queue = queue[1:]

		// stopping condition for search: found a wall
		if startCell.value == 9 {
			continue
		} else {
			(*data)[startCell.x][startCell.y].basin = basinIdx
		}

		neighborCoords := [][]int{
			{startCell.x - 1, startCell.y},
			{startCell.x + 1, startCell.y},
			{startCell.x, startCell.y - 1},
			{startCell.x, startCell.y + 1},
		}
		for _, coordPair := range neighborCoords {
			if isValidCell(visited, coordPair[0], coordPair[1], sizeX, sizeY) {
				queue = append(queue, (*data)[coordPair[0]][coordPair[1]])
				visited[coordPair[0]][coordPair[1]] = true
			}
		}
	}
}

func isValidCell(visited [][]bool, idx int, idy int, sizeX int, sizeY int) bool {
	if idx < 0 || idx >= sizeX {
		return false
	}
	if idy < 0 || idy >= sizeY {
		return false
	}
	if visited[idx][idy] {
		return false
	}
	return true
}

func convertIntArrayToCellArray(content [][]int) [][]Cell {
	data := make([][]Cell, len(content))
	for c := range content {
		data[c] = make([]Cell, len(content[c]))
	}

	for x, row := range content {
		for y, val := range row {
			data[x][y] = Cell{
				x: x, y: y, value: val, isLowPoint: false, basin: 0,
			}
		}
	}
	return data
}

// Returns a list of cells that are their local minima in X+Y
// thoguht about doing an interface here but fundamentally doing a different
// problem. if I refactored part 1 to the same style of
//   1) iterate+mark points
//   2) iterate+calcultae
// then this could interface well...
func findLowPointsOfCells(data *[][]Cell) {
	// assume rectangular
	sizeX := len(*data)
	sizeY := len((*data)[0])

	for idx, row := range *data {
		for idy, cell := range row {
			if (idx-1) >= 0 && (cell.value >= (*data)[idx-1][idy].value) {
				continue
			} else if (idx+1) < sizeX && (cell.value >= (*data)[idx+1][idy].value) {
				continue
			} else if (idy-1) >= 0 && (cell.value >= (*data)[idx][idy-1].value) {
				continue
			} else if (idy+1) < sizeY && (cell.value >= (*data)[idx][idy+1].value) {
				continue
			} else {
				(*data)[idx][idy].isLowPoint = true
			}
		}
	}
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
