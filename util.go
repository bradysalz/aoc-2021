package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Files error out a lot, so make an error handler warpper
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileToIntArray(path string) [][]int {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	retVal := make([][]int, 0)
	for l, line := range strings.Split(string(lines), "\n") {
		retVal = append(retVal, make([]int, 0))
		for _, word := range strings.Split(line, "") {
			val, err := strconv.ParseInt(word, 0, 0)
			if err != nil {
				panic(err)
			}
			retVal[l] = append(retVal[l], int(val))
		}
	}
	return retVal
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

func removeFromStringSlice(s []string, valToRemove string) []string {
	for i, val := range s {
		if val == valToRemove {
			// Swap-n-drop
			s[i] = s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	log.Fatal("Could not find the value in slice")
	return s
}

// Return true if a.sort == b.sort, else false
func checkUnorderedStringEqual(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return subStringInString(a, b)
}

// Return true if a \in b, else false
func subStringInString(a string, b string) bool {
	exists := make(map[rune]bool)
	for _, value := range b {
		exists[value] = true
	}
	for _, value := range a {
		if !exists[value] {
			return false
		}
	}
	return true
}

// Return a | b for strings
func stringUnion(a string, b string) string {
	exists := make(map[rune]bool)
	for _, value := range a {
		exists[value] = true
	}
	for _, value := range b {
		exists[value] = true
	}

	retVal := ""
	for key, val := range exists {
		if val {
			retVal += string(key)
		}
	}
	return retVal
}

// Return a \not \in b (a should be larger than b)
func stringIntersection(a string, b string) string {
	exists := make(map[rune]bool)
	for _, value := range a {
		exists[value] = true
	}
	for _, value := range b {
		exists[value] = false
	}
	retVal := ""
	for key, val := range exists {
		if val {
			retVal += string(key)
		}
	}
	return retVal
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

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
