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

func getChecksum(r []int) int {
	result := 0
	for idx, val := range r {
		if val >= 0 {
			result += idx * val
		}
	}
	return result
}

func part1(diskMap []int) int {
	r := getRepresentation(diskMap)
	for moveBlocks(r) {
	}
	return getChecksum(r)
}

func getNextFile(startIdx int, r []int) (int, int) {
	for ; startIdx >= 0 && r[startIdx] < 0; startIdx-- {
	}
	if startIdx >= 0 {
		var i int
		for i = startIdx; i >= 0 && r[i] == r[startIdx]; i-- {
		}
		i++
		if r[i] == r[startIdx] {
			return i, startIdx
		}
	}
	return -1, -1
}

func moveFile(start, end int, r []int) {
	var gapStart, gapEnd int
	for gapStart = 0; gapStart < start; gapStart++ {
		if r[gapStart] < 0 {
			for gapEnd = gapStart + 1; gapEnd < len(r) && r[gapEnd] < 0; gapEnd++ {
			}
			if gapEnd-gapStart >= end-start+1 {
				for i := 0; i < end-start+1; i++ {
					r[gapStart+i] = r[start+i]
					r[start+i] = -1
				}
				return
			}
		}
	}
}

func printRepr(r []int) {
	s := ""
	for _, el := range r {
		if el < 0 {
			s += "."
		} else {
			s += strconv.Itoa(el)
		}
	}
	fmt.Println(s)
}

func part2(diskMap []int) int {
	r := getRepresentation(diskMap)
	for start, end := getNextFile(len(r)-1, r); start >= 0 && end >= 0; start, end = getNextFile(start-1, r) {
		moveFile(start, end, r)
	}
	return getChecksum(r)
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
	fmt.Printf("PART 2: %v\n", part2(input))
}
