package main

import (
	"io/ioutil"
	"strings"
	"unicode"
)

func dayTwelvePartOne() int {
	pathways := createPathways("input/day_12.txt")
	paths := findAllPaths(pathways, false)
	return len(paths)
}

func dayTwelvePartTwo() int {
	pathways := createPathways("input/day_12.txt")
	paths := findAllPaths(pathways, true)
	return len(paths)
}

// Run a breadth first search to find all the pathways
func findAllPaths(pathways map[string][]string, isPartTwo bool) [][]string {
	paths := make([][]string, 0) // all paths found so far
	queue := make([][]string, 0) // next sites to explore

	// append starting sites
	for p, path := range pathways["start"] {
		queue = append(queue, make([]string, 0))
		queue[p] = append(queue[p], path)
	}

	for len(queue) > 0 {
		startPath := queue[0]
		queue = queue[1:]

		currCave := startPath[len(startPath)-1]
		for _, nextCave := range pathways[currCave] {
			if nextCave == "start" {
				continue
			}

			if nextCave == "end" {
				paths = append(paths, startPath)
				continue
			}

			startStr := ""
			for _, s := range startPath {
				startStr += s
			}
			if !isPartTwo {
				// PART 1 SOLVER
				// if already visited that lower-case cave, skip it
				if unicode.IsLower(rune(nextCave[0])) && strings.Contains(startStr, nextCave) {
					// this works because caves names are either all lower case or all upper case
					continue
				}
			} else {
				// PART 2 SOLVER
				doubleCave := ""
				smallCaves := make(map[string]int, 0)
				for _, s := range startPath {
					if unicode.IsLower(rune(s[0])) {
						smallCaves[s] += 1
						if smallCaves[s] == 2 {
							doubleCave = s
						}
					}
				}
				if doubleCave == "" {
					// can add to path, have not visited a small cave twice
				} else if unicode.IsLower(rune(nextCave[0])) && strings.Contains(startStr, nextCave) {
					continue
				}
			}

			nextPath := make([]string, 0)
			for s := range startPath {
				nextPath = append(nextPath, startPath[s])
			}
			nextPath = append(nextPath, nextCave)
			queue = append(queue, nextPath)
		}
	}
	return paths
}

func createPathways(path string) map[string][]string {
	content, err := ioutil.ReadFile(path)
	check(err)

	pathways := make(map[string][]string, 0)
	data := string(content)
	for _, row := range strings.Split(data, "\n") {
		words := strings.Split(row, "-")
		from, to := words[0], words[1]

		// golang doesn't have sets =/
		_, ok := pathways[from]
		if !ok {
			pathways[from] = make([]string, 0)
		}
		foundVal := false
		for _, s := range pathways[from] {
			if s == to {
				foundVal = true
				break
			}
		}
		if !foundVal {
			pathways[from] = append(pathways[from], to)
		}

		// add reverse path
		_, ok = pathways[to]
		if !ok {
			pathways[to] = make([]string, 0)
		}
		foundVal = false
		for _, s := range pathways[to] {
			if s == from {
				foundVal = true
				break
			}
		}
		if !foundVal {
			pathways[to] = append(pathways[to], from)
		}
	}
	return pathways
}
