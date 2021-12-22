package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Forward string = "forward"
	Down    string = "down"
	Up      string = "up"
)

func dayTwoPartOne() int {
	file, err := os.Open("input/day_02.txt")
	check(err)
	defer file.Close()

	var horizPosition int
	var depth int
	var val int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		val, err = strconv.Atoi(words[1])
		check(err)
		switch words[0] {
		case Forward:
			horizPosition += val
		case Down:
			depth += val
		case Up:
			depth -= val
		default:
			fmt.Println("Unexpected value here")
		}
	}
	return horizPosition * depth
}

func dayTwoPartTwo() int {
	file, err := os.Open("input/day_02.txt")
	check(err)
	defer file.Close()

	var horizPosition int
	var depth int
	var aim int
	var val int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		val, err = strconv.Atoi(words[1])
		check(err)
		switch words[0] {
		case Forward:
			horizPosition += val
			depth += (aim * val)
		case Down:
			aim += val
		case Up:
			aim -= val
		default:
			fmt.Println("Unexpected value here")
		}
	}
	return horizPosition * depth
}
