package main

import (
	"bufio"
	"log"
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

func getLengthOfInputWord(path string) (int, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	// This is probably super inefficient but it works, shrug
	var val string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val = scanner.Text()
		break
	}
	return len(val), scanner.Err()
}

func readFileToNumericArray(path string) ([]int, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var vals []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.ParseInt(scanner.Text(), 2, 32)
		check(err)
		vals = append(vals, int(val))
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

// Remove the idx-th element from the slice s
func removeFromSlice(s []int, idx int) []int {
	// Swap the last element and the "to be removed" one
	s[idx] = s[len(s)-1]
	// Drop the last element
	return s[:len(s)-1]
}

// Check if the pos-th bit is set to 1
func hasBitSet(n int, pos int) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// Get the maximum value and where it (first) occurs
func getMaxValueAndIndex(arr []int) (int, int) {
	if len(arr) == 0 {
		log.Fatal("no maximum of empty array")
	}

	maxIdx := 0
	maxValue := arr[maxIdx]

	for a := range arr {
		if arr[a] > maxValue {
			maxValue = arr[a]
			maxIdx = a
		}
	}
	return maxValue, maxIdx
}
