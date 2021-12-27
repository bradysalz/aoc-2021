package main

import (
	"bufio"
	"fmt"
	"os"
)

func dayFourteenPartOne() int {
	seq, seqMap, last := readInputDayFourteen("input/day_14.txt")
	scoreSequence(seq, last)
	max, min := 0, 0
	for s := 0; s < 10; s++ {
		seq, last = updateSequence(seq, seqMap, last)
		max, min = scoreSequence(seq, last)
	}

	return max - min
}
func dayFourteenPartTwo() int {
	seq, seqMap, last := readInputDayFourteen("input/day_14.txt")
	max, min := 0, 0
	for s := 0; s < 40; s++ {
		seq, last = updateSequence(seq, seqMap, last)
		max, min = scoreSequence(seq, last)
	}

	return max - min
}

func scoreSequence(seq map[string]int, last string) (int, int) {
	scoreMap := make(map[string]int, 0)

	for k, v := range seq {
		scoreMap[string(k[0])] += v
	}
	scoreMap[string(last[0])] += 1
	scoreMap[string(last[1])] += 1

	// init with a random letter
	minVal := scoreMap["N"]
	maxVal := minVal
	for _, val := range scoreMap {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, minVal
}

func updateSequence(seq map[string]int, seqMap map[string]string, last string) (map[string]int, string) {
	out := make(map[string]int, 0)

	for k, v := range seq {
		c := seqMap[k]
		out[string(k[0])+c] += v
		out[c+string(k[1])] += v

	}

	temp := seqMap[last]
	newSecondLast := string(last[0]) + temp
	newLast := temp + string(last[1])
	out[newSecondLast] += 1

	return out, newLast
}

func readInputDayFourteen(path string) (map[string]int, map[string]string, string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// la la inefficient idc
	sequenceFound := false
	newlineFound := false
	sequence := make(map[string]int, 0)
	seqMap := make(map[string]string, 0)
	lastPair := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !sequenceFound {
			tmp := scanner.Text()
			for t := 0; t < len(tmp)-2; t++ {
				sequence[string(tmp[t])+string(tmp[t+1])] += 1
			}
			sequenceFound = true
			lastPair = string(tmp[len(tmp)-2]) + string(tmp[len(tmp)-1])
			continue
		}
		if !newlineFound {
			scanner.Text()
			newlineFound = true
			continue
		}

		var key string
		var val string
		_, err := fmt.Sscanf(scanner.Text(), "%s -> %s\n", &key, &val)
		if err != nil {
			panic(err)
		}
		seqMap[key] = val
	}
	return sequence, seqMap, lastPair
}
