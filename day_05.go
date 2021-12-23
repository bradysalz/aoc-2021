package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func dayFivePartOne(isPartTwo bool) int {
	data := readInputDayFive("input/day_05.txt")

	// init the board
	boardSize := findLargestElement(data) + 1
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	// build the map
	for l, line := range data {
		if line.start.x == line.end.x {
			// Re-organized to make the for loop easy
			if line.start.y > line.end.y {
				tmp := data[l].start.y
				data[l].start.y = data[l].end.y
				data[l].end.y = tmp
			}
			for idx := data[l].start.y; idx <= data[l].end.y; idx++ {
				board[line.start.x][idx] += 1
			}
		} else if line.start.y == line.end.y {
			// Re-organized to make the for loop easy
			if line.start.x > line.end.x {
				tmp := data[l].start.x
				data[l].start.x = data[l].end.x
				data[l].end.x = tmp
			}
			for idx := data[l].start.x; idx <= data[l].end.x; idx++ {
				board[idx][line.start.y] += 1
			}
		} else if isPartTwo {
			slope := (line.end.y - line.start.y) / (line.end.x - line.start.x)
			if !(slope == 1 || slope == -1) {
				// turns out this doesn't exist. all inputs are valid.
				continue
			}
			var incX, incY int
			if line.start.x > line.end.x {
				incX = -1
			} else {
				incX = 1
			}
			if line.start.y > line.end.y {
				incY = -1
			} else {
				incY = 1
			}
			// 45 deg means equal steps both ways, only need one loop
			idy := line.start.y
			for idx := data[l].start.x; idx != (data[l].end.x + incX); idx += incX {
				board[idx][idy] += 1
				idy += incY
			}
		}
	}

	numOverlap := 0
	for r, row := range board {
		for c := range row {
			if board[r][c] > 1 {
				numOverlap += 1
			}
		}
	}
	return numOverlap
}

// Find the largest x or y coordinate in the dataset
// Used to sized the board array
func findLargestElement(data []Line) int {
	maxValue := 0
	for _, line := range data {
		if line.start.x > maxValue {
			maxValue = line.start.x
		}
		if line.start.y > maxValue {
			maxValue = line.start.y
		}
		if line.end.x > maxValue {
			maxValue = line.end.x
		}
		if line.end.y > maxValue {
			maxValue = line.end.y
		}
	}
	return maxValue
}

func readInputDayFive(path string) []Line {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	retVal := make([]Line, 0)
	for _, line := range strings.Split(string(lines), "\n") {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(string(line), "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(err)
		}

		retVal = append(retVal, Line{
			start: Point{x: x1, y: y1},
			end:   Point{x: x2, y: y2},
		})
	}
	return retVal
}
