package main

import (
	"bufio"
	"os"
	"strconv"
)

// Files error out a lot, so make an error handler warpper
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileToStringArray(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readFileToNumericArray(path string) ([]int, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var vals []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		vals = append(vals, val)
	}
	return vals, scanner.Err()
}

// Calculate the the logical inverse of a value
// Have had issues with sign extensions, so explicitly require a flag length
func calculateInverseValue(value int, flagLength int) int {
	var flag int
	for i := 0; i < flagLength; i++ {
		flag += 1 << i
	}
	// reminder: golang can't do !, so we have to do xor FF instead
	// a xor !a = 1   ->   1 xor a = !a
	return (value ^ flag)
}
