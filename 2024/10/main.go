package main

import (
	"bufio"
	"fmt"
	"os"
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

func getStartPoints(data [][]int) [][2]int {
	starts := [][2]int{}
	for x, row := range data {
		for y, val := range row {
			if val == 0 {
				starts = append(starts, [2]int{x, y})
			}
		}
	}
	return starts
}

func dfs(data [][]int, x, y int, stop func(val, x, y int) bool) {
	if !stop(data[x][y], x, y) {
		for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			xx := x + d[0]
			yy := y + d[1]
			if xx >= 0 && xx < len(data) &&
				yy >= 0 && yy < len(data[0]) &&
				data[xx][yy]-data[x][y] == 1 {
				dfs(data, xx, yy, stop)
			}
		}
	}
}

func part1(data [][]int) int {
	starts := getStartPoints(data)
	result := 0
	for _, start := range starts {
		visited := make(map[[2]int]bool)
		stop := func(val, x, y int) bool {
			shouldStop := val == 9
			pos := [2]int{x, y}
			_, ok := visited[pos]
			if shouldStop && !ok {
				result++
				visited[pos] = true
			}
			return shouldStop
		}
		dfs(data, start[0], start[1], stop)
	}
	return result
}

func part2(data [][]int) int {
	starts := getStartPoints(data)
	result := 0
	stop := func(val, x, y int) bool {
		if val == 9 {
			result++
			return true
		}
		return false
	}
	for _, start := range starts {
		dfs(data, start[0], start[1], stop)
	}
	return result
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
	fmt.Printf("PART 2: %v\n", part2(input))
}
