package main

import (
	"errors"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type Sequence struct {
	in  []string
	out []string
}

// Converts a number to it's number of segments
var characterLengthMap = map[int]int{
	0: 6,
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

func dayEightPartOne() int {
	sequences := readInputDayEight("input/day_08.txt")
	totalFound := 0
	partOneSequence := []int{2, 3, 4, 7}
	for _, seq := range sequences {
		for _, word := range seq.out {
			if inRangeDayEight(word, partOneSequence) {
				totalFound += 1
			}
		}
	}
	return totalFound
}

func inRangeDayEight(a string, vals []int) bool {
	for v := range vals {
		if len(a) == vals[v] {
			return true
		}
	}
	return false
}

func dayEightPartTwo() int {
	sequences := readInputDayEight("input/day_08.txt")

	cumSum := 0
	for _, sequence := range sequences {
		characterMap := solveAllNumberMapping(sequence.in)
		rowSum, err := calculateOutputSequence(sequence.out, characterMap)
		if err != nil {
			panic(err)
		}
		cumSum += rowSum
	}
	return cumSum
}

func calculateOutputSequence(seq []string, characterMap map[int]string) (int, error) {
	rowSum := 0
	for s, word := range seq {
		for i := 0; i < 11; i++ {
			if checkUnorderedStringEqual(word, characterMap[i]) {
				rowSum += i * int(math.Pow10(3-s))
				break
			}
			if i == 10 {
				return 0, errors.New("Could not find a matching number")
			}
		}
	}
	return rowSum, nil
}

func solveAllNumberMapping(data []string) map[int]string {
	characterMap := make(map[int]string)
	temp := data
	var err error

	// Solve easy ones (1,4,7,8)
	characterMap[1] = solveEasyMapping(data, 1)
	temp = removeFromStringSlice(temp, characterMap[1])

	characterMap[4] = solveEasyMapping(data, 4)
	temp = removeFromStringSlice(temp, characterMap[4])

	characterMap[7] = solveEasyMapping(data, 7)
	temp = removeFromStringSlice(temp, characterMap[7])

	characterMap[8] = solveEasyMapping(data, 8)
	temp = removeFromStringSlice(temp, characterMap[8])

	// 2, 3, 5 are all same length
	characterMap[3], err = solveForThree(temp, characterMap[1])
	check(err)
	temp = removeFromStringSlice(temp, characterMap[3])

	characterMap[5], err = solveForFive(temp, characterMap[4], characterMap[1])
	check(err)
	temp = removeFromStringSlice(temp, characterMap[5])

	characterMap[2], err = solveForTwo(temp)
	check(err)
	temp = removeFromStringSlice(temp, characterMap[2])

	// 0, 6, 9
	characterMap[9], err = solveForNine(temp, characterMap[4], characterMap[7])
	check(err)
	temp = removeFromStringSlice(temp, characterMap[9])

	characterMap[6], err = solveForSix(temp, characterMap[5])
	check(err)
	temp = removeFromStringSlice(temp, characterMap[6])

	if len(temp) != 1 {
		panic(errors.New("uh oh list too long"))
	}
	characterMap[0] = temp[0]

	return characterMap
}

func solveEasyMapping(data []string, val int) string {
	for _, word := range data {
		if len(word) == characterLengthMap[val] {
			return word
		}
	}
	log.Fatal("Could not find matching word")
	return ""
}

// Critera for finding 6
// - 5 \in 6
// - only have 0, 6 left (8+9 (and 5, duh) match)
func solveForSix(data []string, fiveSegment string) (string, error) {
	for _, word := range data {
		if subStringInString(fiveSegment, word) {
			return word, nil
		}
	}
	return "", errors.New("Could not find 0")
}

// Critera for finding 9
// - (4+7) \in 9
// don't check length, all vals are len(6) now
func solveForNine(data []string, fourSegment string, sevenSegment string) (string, error) {
	combinedSegment := stringUnion(fourSegment, sevenSegment)
	for _, word := range data {
		if subStringInString(combinedSegment, word) {
			return word, nil
		}
	}
	return "", errors.New("Could not find 9")
}

// Criteria for finding 2
// - 2 has 5 letters
// - should be the only one left
func solveForTwo(data []string) (string, error) {
	for _, word := range data {
		if len(word) == 5 {
			return word, nil
		}
	}
	return "", errors.New("Could not find a 2")
}

// Criteria for finding 5
// - 5 has the ("4" - "1") characters
// - 5 has length 5
func solveForFive(data []string, fourSegments string, oneSegments string) (string, error) {
	segmentIntersection := stringIntersection(fourSegments, oneSegments)
	for _, word := range data {
		if len(word) == 5 && subStringInString(segmentIntersection, word) {
			return word, nil
		}
	}
	return "", errors.New("Could not find a 5")

}

// Criteria for finding three:
// - 3 has the "1" characters
// - 3 has length 5
func solveForThree(data []string, oneSegments string) (string, error) {
	for _, word := range data {
		if len(word) == 5 && subStringInString(oneSegments, word) {
			return word, nil
		}
	}
	return "", errors.New("Could not find a 3")
}

// Turns the input data into array of Sequences
func readInputDayEight(path string) []Sequence {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	sequences := make([]Sequence, 0)
	for _, line := range lines {
		inputSeq := make([]string, 0)
		outputSeq := make([]string, 0)
		for l, half := range strings.Split(string(line), " | ") {
			for _, word := range strings.Split(half, " ") {
				if l == 0 {
					inputSeq = append(inputSeq, word)
				} else {
					outputSeq = append(outputSeq, word)
				}
			}
		}
		sequences = append(sequences, Sequence{in: inputSeq, out: outputSeq})
	}
	return sequences
}
