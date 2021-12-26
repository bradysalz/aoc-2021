package main

import (
	"sort"
	"strings"
)

const OPEN_CHARS = "({[<"

var CHAR_SCORE_P1 = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var CHAR_SCORE_P2 = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var OPEN_CLOSE_MAP = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func dayTenPartOne() int {
	lines, err := readFileToStringArray("input/day_10.txt")
	check(err)

	total := 0
	for _, line := range lines {
		lastScore := checkLineAndScorePartOne(line)
		total += lastScore
	}
	return total
}

func dayTenPartTwo() int {
	lines, err := readFileToStringArray("input/day_10.txt")
	check(err)

	total := make([]int, 0)
	for _, line := range lines {
		lastScore := checkLineAndScorePartTwo(line)
		if lastScore == 0 {
			continue
		}
		total = append(total, lastScore)
	}
	sort.Ints(total)
	return total[len(total)/2]
}

// Return 0 if line is incomplete, else return the score of the wrong character
func checkLineAndScorePartOne(line string) int {
	openingChars := ""

	for _, c := range line {
		s := string(c)
		lastMatch := OPEN_CLOSE_MAP[getLastChar(openingChars)]
		if strings.Contains(OPEN_CHARS, s) {
			openingChars += s
		} else if s == lastMatch {
			// found a match, pop off last character
			if len(openingChars) < 1 {
				openingChars = ""
			} else {
				openingChars = openingChars[:len(openingChars)-1]
			}
		} else {
			return CHAR_SCORE_P1[s]
		}
	}
	// Incomplete line
	return 0
}

// Return 0 if line is corrupt, else return the score of the missing characters
func checkLineAndScorePartTwo(line string) int {
	openingChars := ""

	for _, c := range line {
		s := string(c)
		lastMatch := OPEN_CLOSE_MAP[getLastChar(openingChars)]
		if strings.Contains(OPEN_CHARS, s) {
			openingChars += s
		} else if s == lastMatch {
			// found a match, pop off last character
			if len(openingChars) < 1 {
				openingChars = ""
			} else {
				openingChars = openingChars[:len(openingChars)-1]
			}
		} else {
			// corrupt line
			return 0
		}
	}

	score := 0
	for _, c := range Reverse(openingChars) {
		score *= 5
		s := string(c)
		score += CHAR_SCORE_P2[OPEN_CLOSE_MAP[s]]

	}
	return score
}

func getLastChar(s string) string {
	if len(s) > 0 {
		return string(s[len(s)-1])
	} else {
		return ""
	}
}
