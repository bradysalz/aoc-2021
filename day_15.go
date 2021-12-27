package main

import (
	"math"
)

type vertex struct {
	x        int
	y        int
	risk     int
	pathRisk int
}

func dayFifteenPartOne(isPartTwo bool) int {
	data := readFileToIntArray("input/day_15.txt")
	graph := convertToGraph(data)
	if isPartTwo {
		graph = biggifyGraph(graph)
	}
	solveForMinCost(&graph)
	return graph[len(graph)-1][len(graph[0])-1].pathRisk
}

func biggifyGraph(graph [][]vertex) [][]vertex {
	lenX, lenY := len(graph), len(graph[0])
	newGraph := make([][]vertex, 5*lenX)
	for x := 0; x < 5*lenY; x++ {
		newGraph[x] = make([]vertex, 5*lenY)
	}

	for x := range newGraph {
		for y := range newGraph[0] {
			newCell := vertex{
				x:        x,
				y:        y,
				risk:     graph[x%lenX][y%lenY].risk + int(x/lenX) + int(y/lenY),
				pathRisk: math.MaxInt64,
			}
			if newCell.risk > 9 {
				newCell.risk = (newCell.risk % 10) + 1
			}
			newGraph[newCell.x][newCell.y] = newCell
		}
	}
	return newGraph
}

func solveForMinCost(graph *[][]vertex) {
	(*graph)[0][0].pathRisk = 0
	queue := make([]vertex, 0)
	queue = append(queue, (*graph)[0][0])

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		offsets := [][]int{
			{1, 0},
			{0, 1},
			{-1, 0}, // need to lurk on reddit for help here, didn't think to look up/left
			{0, -1},
		}

		for _, offset := range offsets {
			if v.x+offset[0] >= len(*graph) || v.y+offset[1] >= len((*graph)[0]) {
				continue
			}

			if v.x+offset[0] < 0 || v.y+offset[1] < 0 {
				continue
			}
			currCell := (*graph)[v.x][v.y]
			nextCell := (*graph)[v.x+offset[0]][v.y+offset[1]]

			if nextCell.pathRisk > currCell.pathRisk+nextCell.risk {
				(*graph)[v.x+offset[0]][v.y+offset[1]].pathRisk = currCell.pathRisk + nextCell.risk
				queue = append(queue, nextCell)
			}
		}
	}
}

func convertToGraph(data [][]int) (graph [][]vertex) {
	graph = make([][]vertex, len(data))
	for x := range data {
		graph[x] = make([]vertex, len(data[0]))
		for y := range data[x] {
			graph[x][y] = vertex{x: x, y: y, risk: data[x][y], pathRisk: math.MaxInt64}
		}
	}
	return graph
}
