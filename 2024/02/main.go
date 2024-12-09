package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput(filepath string) ([][]int, error) {
	result := [][]int{}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineRaw := scanner.Text()
		line := []int{}
		for _, numberRaw := range strings.Fields(lineRaw) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				return nil, err
			}
			line = append(line, number)
		}
		result = append(result, line)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func part1(input [][]int) int {
	result := 0
	for _, level := range input {
		isSafe := true
		for idx := 1; idx < len(level)-1; idx++ {
			diff1 := level[idx] - level[idx-1]
			diff2 := level[idx+1] - level[idx]
			if diff1*diff2 <= 0 || math.Abs(float64(diff1)) > 3 || math.Abs(float64(diff2)) > 3 {
				isSafe = false
			}
		}
		if isSafe {
			result++
		}
	}
	return result
}

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	p1 := part1(input)
	fmt.Printf("PART 1: %v\n", p1)
}
