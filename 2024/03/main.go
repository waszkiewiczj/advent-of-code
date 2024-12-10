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

func part1(input string) int {
	multiplies := [][]int{}
	for idx := 0; idx < len(input); idx++ {
		if idx+4 < len(input) && input[idx:idx+4] == "mul(" {
			newIdx := idx + 4
			for ; newIdx < len(input) && string(input[newIdx]) != ")"; newIdx++ {
			}
			if newIdx < len(input) {
				argsRaw := input[idx+4 : newIdx]
				numbersRaw := strings.Split(argsRaw, ",")
				if len(numbersRaw) == 2 {
					n1, err1 := strconv.Atoi(numbersRaw[0])
					n2, err2 := strconv.Atoi(numbersRaw[1])
					if err1 == nil && err2 == nil {
						multiplies = append(multiplies, []int{n1, n2})
					}
				}
			}
		}
	}
	result := 0
	for _, pair := range multiplies {
		result += pair[0] * pair[1]
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
