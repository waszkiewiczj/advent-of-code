package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filepath string) (string, error) {
	result, err := os.ReadFile(filepath)
	return string(result), err
}

func parseMulArgs(argsRaw string) ([]int, error) {
	numbersRaw := strings.Split(argsRaw, ",")
	if len(numbersRaw) != 2 {
		return nil, fmt.Errorf("Invalid number of args")
	}
	n1, err := strconv.Atoi(numbersRaw[0])
	if err != nil {
		return nil, err
	}
	n2, err := strconv.Atoi(numbersRaw[1])
	if err != nil {
		return nil, err
	}
	return []int{n1, n2}, nil
}

func parseMul(s string, idx int) ([]int, error) {
	if idx+4 < len(s) && s[idx:idx+4] == "mul(" {
		newIdx := idx + 4
		for ; newIdx < len(s) && string(s[newIdx]) != ")"; newIdx++ {
		}
		if newIdx < len(s) {
			argsRaw := s[idx+4 : newIdx]
			return parseMulArgs(argsRaw)
		}
	}
	return nil, fmt.Errorf("Mul operation not found")
}

func parseDo(s string, idx int) (bool, error) {
	if idx+4 < len(s) && s[idx:idx+4] == "do()" {
		return true, nil
	}
	if idx+7 < len(s) && s[idx:idx+7] == "don't()" {
		return false, nil
	}
	return false, fmt.Errorf("Do/Don't operation not found")
}

func part1(input string) int {
	result := 0
	for idx := 0; idx < len(input); idx++ {
		numbers, err := parseMul(input, idx)
		if err == nil {
			result += numbers[0] * numbers[1]
		}
	}
	return result
}

func part2(input string) int {
	result := 0
	include := true
	for idx := 0; idx < len(input); idx++ {
		newInclude, err := parseDo(input, idx)
		if err == nil {
			include = newInclude
		}
		numbers, err := parseMul(input, idx)
		if err == nil && include {
			result += numbers[0] * numbers[1]
		}
	}
	return result
}
func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PART 1: %v\n", part1(input))
	fmt.Printf("PART 2: %v\n", part2(input))
}
