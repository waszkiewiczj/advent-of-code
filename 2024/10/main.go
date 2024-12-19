package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dominikbraun/graph"
)

func parseInput(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	result := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int{}
		lineRaw := scanner.Text()
		for _, numberRaw := range lineRaw {
			line = append(line, int(numberRaw-'0'))
		}
		result = append(result, line)
	}
	return result
}

func iterateOverElements(data [][]int, f func(x, y, val int)) {
	for x, row := range data {
		for y, val := range row {
			f(x, y, val)
		}
	}
}

func getVertexLabel(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func initGraph(data [][]int) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Directed())

	iterateOverElements(data, func(x, y, val int) {
		label := getVertexLabel(x, y)
		g.AddVertex(label)
	})

	iterateOverElements(data, func(x, y, val int) {
		label := getVertexLabel(x, y)
		for _, d := range [][]int{
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		} {
			xx := x + d[0]
			yy := y + d[1]
			if xx >= 0 && yy >= 0 && xx < len(data) && yy < len(data[0]) && data[xx][yy]-val == 1 {
				g.AddEdge(label, fmt.Sprintf("%v-%v", xx, yy))
			}
		}
	})
	return g
}

func part1(data [][]int) int {
	g := initGraph(data)
	starts := [][]int{}
	ends := [][]int{}
	iterateOverElements(data, func(x, y, val int) {
		if val == 0 {
			starts = append(starts, []int{x, y})
		}
		if val == 9 {
			ends = append(ends, []int{x, y})
		}
	})

	result := 0
	for _, start := range starts {
		startLabel := fmt.Sprintf("%v-%v", start[0], start[1])
		for _, end := range ends {
			endLabel := fmt.Sprintf("%v-%v", end[0], end[1])
			_, err := graph.ShortestPath(g, startLabel, endLabel)
			if err == nil {
				result++
			}
		}
	}
	return result
}

func dfs(data [][]int, x, y int, result *int) {
	if data[x][y] == 9 {
		*result++
	} else {
		for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			xx := x + d[0]
			yy := y + d[1]
			if xx >= 0 && xx < len(data) &&
				yy >= 0 && yy < len(data[0]) &&
				data[xx][yy]-data[x][y] == 1 {
				dfs(data, xx, yy, result)
			}
		}
	}
}

func part2(data [][]int) int {
	starts := [][]int{}
	iterateOverElements(data, func(x, y, val int) {
		if val == 0 {
			starts = append(starts, []int{x, y})
		}
	})
	result := 0
	for _, start := range starts {
		dfs(data, start[0], start[1], &result)
	}
	return result
}

func main() {
	input := parseInput("input.txt")

	// fmt.Printf("PART 1: %v\n", part1(input))
	fmt.Printf("PART 2: %v\n", part2(input))
}
