package main

import (
	"bufio"
	"os"
	"strconv"
)

func dayOnePartOne() uint {
	file, err := os.Open("input/day_01.txt")
	check(err)
	defer file.Close()

	var bFirstElement bool = true
	var currVal int
	var lastVal int
	var increaseCount uint = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if bFirstElement {
			currVal, err = strconv.Atoi(scanner.Text())
			check(err)
			bFirstElement = false
		}
		lastVal = currVal
		currVal, err = strconv.Atoi(scanner.Text())
		check(err)

		if currVal > lastVal {
			increaseCount += 1
		}
	}

	return increaseCount
}

func dayOnePartTwo() uint {
	file, err := os.Open("input/day_01.txt")
	check(err)
	defer file.Close()

	readValues := make([]int, 0)
	var increaseCount uint = 0
	var val int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err = strconv.Atoi(scanner.Text())
		check(err)
		readValues = append(readValues, val)

		if len(readValues) < 4 {
			// Wait until we have enough samples
			continue
		}

		if readValues[3] > readValues[0] {
			increaseCount += 1
		}

		readValues = readValues[1:]

	}
	return increaseCount

}
