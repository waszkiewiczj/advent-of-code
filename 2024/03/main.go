package main

import (
	"fmt"
	"os"
)

func parseInput(filepath string) (string, error) {
	result, err := os.ReadFile(filepath)
	return string(result), err
}

func part1(input string) int {
	for idx := 0; idx < len(input); idx++ {
		if idx+4 < len(input) && input[idx:idx+4] == "mul(" {
			newIdx := idx + 4
			for ; newIdx < len(input) && string(input[newIdx]) != ")"; newIdx++ {
			}
			if newIdx < len(input) {
				numbersRaw := input[idx+4 : newIdx]
				fmt.Println(numbersRaw)
			}
		}
	}
	return 0
}

func main() {
	input, err := parseInput("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PART 1: %v\n", part1(input))
}
