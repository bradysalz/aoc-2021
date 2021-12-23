package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	value  int
	called bool
}

const LEN_ROW_OR_COL = 5

func dayFourPartOne() int {
	draws, boards, err := readInput("input/day_04.txt")
	if err != nil {
		panic(err)
	}

	winningBoards := make([]int, 0)
	lastDrawn := 0
	for _, drawnValue := range draws {
		// if we found a winner last time, we're good
		if len(winningBoards) > 0 {
			break
		}

		// Update all boards
		for b, board := range boards {
			for r, row := range board {
				for c, cell := range row {
					if cell.value == drawnValue {
						boards[b][r][c].called = true
					}
				}
			}
		}

		// Check for winners
		for b, board := range boards {
			if checkIfBoardWins(board) {
				winningBoards = append(winningBoards, b)
				lastDrawn = drawnValue
			}
		}
	}

	uncalledSum := 0
	for _, row := range boards[winningBoards[0]] {
		for _, cell := range row {
			if !cell.called {
				uncalledSum += cell.value
			}
		}
	}
	return uncalledSum * lastDrawn
}

// Find the last winning board
func dayFourPartTwo() int {
	draws, boards, err := readInput("input/day_04.txt")
	if err != nil {
		panic(err)
	}

	winningBoards := make([]int, 0)
	lastDrawn := 0
	for _, drawnValue := range draws {
		// Find the last winning board
		if len(winningBoards) == len(boards) {
			break
		}

		// Update all boards
		for b, board := range boards {
			for r, row := range board {
				for c, cell := range row {
					if cell.value == drawnValue {
						boards[b][r][c].called = true
					}
				}
			}
		}

		// Check for winners
		for b, board := range boards {
			if checkIfBoardWins(board) {
				alreadyCounted := false

				// it's insane that go doesn't have a "element in slice" function
				// i'm 99% sure this method is necessary, was getting errors without it
				for _, a := range winningBoards {
					if a == b {
						alreadyCounted = true
						break
					}
				}
				if alreadyCounted {
					continue
				}
				winningBoards = append(winningBoards, b)
				lastDrawn = drawnValue
			}
		}
	}

	uncalledSum := 0
	for _, row := range boards[winningBoards[len(boards)-1]] {
		for _, cell := range row {
			if !cell.called {
				uncalledSum += cell.value
			}
		}
	}
	return uncalledSum * lastDrawn
}

func checkIfBoardWins(board [][]Square) bool {
	// Check rows
	for _, row := range board {
		hasWon := true
		for c := 0; c < LEN_ROW_OR_COL; c++ {
			if !row[c].called {
				hasWon = false
				break
			}
		}
		if hasWon {
			return true
		}
	}

	// Check columns
	for columnIndex := 0; columnIndex < LEN_ROW_OR_COL; columnIndex++ {
		hasWon := true
		for _, row := range board {
			if !row[columnIndex].called {
				hasWon = false
				break
			}
		}
		if hasWon {
			return true
		}
	}

	return false

}

func readInput(path string) ([]int, map[int][][]Square, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	hasReadDraws := false
	boardCount := -1 // galaxy brain for later
	x := 0
	draws := make([]int, 0)
	boards := make(map[int][][]Square)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !hasReadDraws {
			draws = readDelimitedRow(scanner.Text(), ",")
			hasReadDraws = true
			continue
		}

		row := scanner.Text()
		if row == "" {
			boardCount += 1 // it's later
			board := make([][]Square, 0)
			boards[boardCount] = board
			x = 0
			continue
		}

		boards[boardCount] = append(boards[boardCount], make([]Square, 0))
		temp := readDelimitedRow(row, " ")
		boards[boardCount][x] = append(boards[boardCount][x], convertRowToSquares(temp)...)
		x += 1
	}
	return draws, boards, scanner.Err()
}

func convertRowToSquares(row []int) []Square {
	temp := make([]Square, 0)
	for _, val := range row {
		newSquare := Square{
			value:  val,
			called: false,
		}
		temp = append(temp, newSquare)
	}
	return temp
}

// Turn a string of delimited ints to an array of ints
func readDelimitedRow(data string, delimiter string) []int {
	var arr []int
	s := strings.Split(data, delimiter)
	for _, char := range s {
		if char == "" {
			continue
		}

		val, err := strconv.ParseInt(char, 0, 0)
		if err != nil {
			panic(err)
		}
		arr = append(arr, int(val))
	}
	return arr
}
