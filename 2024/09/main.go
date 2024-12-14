package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseInput(filepath string) []int {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	result := []int{}
	for _, r := range raw {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}

func getRepresentation(diskMap []int) []int {
	result := []int{}
	for idx, val := range diskMap {
		c := -1
		if idx%2 == 0 {
			c = idx / 2
		}
		for i := 0; i < val; i++ {
			result = append(result, c)
		}
	}
	return result
}

func moveBlocks(r []int) bool {
	var i, j int
	for i = 0; i < len(r) && r[i] >= 0; i++ {
	}
	for j = len(r) - 1; j >= 0 && r[j] == -1; j-- {
	}
	if i < j {
		r[i], r[j] = r[j], r[i]
		return true
	}
	return false
}

func part1(diskMap []int) int {
	repr := getRepresentation(diskMap)
	for moveBlocks(repr) {
	}
	result := 0
	for idx, val := range repr {
		if val < 0 {
			break
		}
		result += idx * val
	}
	return result
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
}
