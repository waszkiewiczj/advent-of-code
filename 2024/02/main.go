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

func isSafe(report []int) bool {
	for idx := 1; idx < len(report)-1; idx++ {
		diff1 := report[idx] - report[idx-1]
		diff2 := report[idx+1] - report[idx]
		if diff1*diff2 <= 0 || math.Abs(float64(diff1)) > 3 || math.Abs(float64(diff2)) > 3 {
			return false
		}
	}
	return true
}

func part1(input [][]int) int {
	result := 0
	for _, report := range input {
		if isSafe(report) {
			result++
		}
	}
	return result
}

func part2(input [][]int) int {
	result := 0
	for _, report := range input {
		safe := isSafe(report)
		for idx := 0; idx < len(report) && !safe; idx++ {
			newReport := []int{}
			newReport = append(newReport, report[:idx]...)
			newReport = append(newReport, report[idx+1:]...)
			safe = isSafe(newReport)
		}
		if safe {
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

	p2 := part2(input)
	fmt.Printf("PART 2: %v\n", p2)
}
