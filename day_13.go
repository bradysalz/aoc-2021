package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type instr struct {
	dir string
	val int
}

func dayThirteenPartOne() int {
	coords, instrs := readInputDayThirteen("input/day_13_test.txt")
	arr := buildArray(coords)
	foldPaper(&arr, instrs[0])
	sum := countDots(arr)
	return sum
}

func dayThirteenPartTwo() {
	coords, instrs := readInputDayThirteen("input/day_13.txt")
	arr := buildArray(coords)
	for i := range instrs {
		foldPaper(&arr, instrs[i])
	}
	debugArr(arr)
}

func debugArr(arr [][]bool) {
	for y := range arr[0] {
		for x := range arr {
			if arr[x][y] {
				fmt.Print("█")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func foldPaper(arr *[][]bool, i instr) {
	if i.dir == "y" {
		for x := range *arr {
			for v := 0; v < i.val; v++ {
				// are you kidding me??? no |= for bools???
				if (2*i.val - v) >= len((*arr)[x]) {
					continue
				}
				(*arr)[x][v] = (*arr)[x][v] || (*arr)[x][2*i.val-v]
			}
		}
		for x := range *arr {
			(*arr)[x] = (*arr)[x][:i.val]
		}
	} else {
		for y := 0; y < len((*arr)[0]); y++ {
			for v := 0; v < i.val; v++ {
				if (2*i.val - v) >= len((*arr)) {
					continue
				}
				(*arr)[v][y] = (*arr)[v][y] || (*arr)[2*i.val-v][y]
			}
		}
		*arr = (*arr)[:i.val]
	}
}

func countDots(arr [][]bool) int {
	sum := 0
	for _, row := range arr {
		for _, cell := range row {
			if cell {
				sum += 1
			}
		}
	}
	return sum
}

func buildArray(coords []coord) [][]bool {
	sizeX, sizeY := 0, 0
	for _, c := range coords {
		if c.x > sizeX {
			sizeX = c.x
		}
		if c.y > sizeY {
			sizeY = c.y
		}
	}

	arr := make([][]bool, sizeX+1)
	for a := range arr {
		arr[a] = make([]bool, sizeY+1)
	}

	for _, c := range coords {
		arr[c.x][c.y] = true
	}

	return arr
}
func readInputDayThirteen(path string) ([]coord, []instr) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readingCoords := true
	coords := make([]coord, 0)
	instrs := make([]instr, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			readingCoords = false
			continue
		}

		if readingCoords {
			pt := strings.Split(scanner.Text(), ",")
			x, err := strconv.ParseInt(pt[0], 0, 0)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseInt(pt[1], 0, 0)
			if err != nil {
				panic(err)
			}
			coords = append(coords, coord{x: int(x), y: int(y)})
		} else {
			halves := strings.Split(scanner.Text(), "=")
			dir := halves[0][len(halves[0])-1]
			val, err := strconv.ParseInt(halves[1], 0, 0)
			if err != nil {
				panic(err)
			}
			instrs = append(instrs, instr{dir: string(dir), val: int(val)})
		}
	}
	return coords, instrs
}
