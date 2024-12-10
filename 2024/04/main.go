package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput(filepath string) ([]string, error) {
	result := []string{}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func isXMAS(s string) bool {
	return s == "XMAS" || s == "SAMX"
}

func isXMASHorizontal(input []string, i, j int) bool {
	if j+4 <= len(input[i]) {
		s := input[i][j : j+4]
		return isXMAS(s)
	}
	return false
}

func isXMASVertical(input []string, i, j int) bool {
	s := ""
	for z := 0; i+z < len(input) && z < 4; z++ {
		s += string(input[i+z][j])
	}
	return isXMAS(s)
}

func isXMASDiagonal(input []string, i, j int) bool {
	s := ""
	for z := 0; i+z < len(input) && j+z < len(input[i+z]) && z < 4; z++ {
		s += string(input[i+z][j+z])
	}
	return isXMAS(s)
}

func isXMASAntiDiagonal(input []string, i, j int) bool {
	s := ""
	for z := 0; i+z < len(input) && j-z >= 0 && z < 4; z++ {
		s += string(input[i+z][j-z])
	}
	return isXMAS(s)
}

func countXMAS(input []string, i, j int) int {
	xmasFuncs := []func(input []string, i, j int) bool{
		isXMASHorizontal,
		isXMASVertical,
		isXMASDiagonal,
		isXMASAntiDiagonal,
	}

	result := 0
	for _, f := range xmasFuncs {
		if f(input, i, j) {
			result++
		}
	}

	return result
}

func part1(input []string) int {
	result := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			count := countXMAS(input, i, j)
			result += count
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
}
