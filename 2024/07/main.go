package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	testValue int
	numbers   []int
}

func (c *Calibration) isTrue() {

}

func parseInput(filepath string) []Calibration {
	result := []Calibration{}
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slices := strings.Split(line, ":")
		testValue, err := strconv.Atoi(slices[0])
		if err != nil {
			panic(err)
		}
		numbers := []int{}
		for _, numberRaw := range strings.Fields(slices[1]) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
		c := Calibration{testValue, numbers}
		result = append(result, c)
	}
	return result
}

func part1(input []Calibration) int {
	result := 0
	for _, c := range input {
		if evaluate(c.testValue, c.numbers, 0) {
			result += c.testValue
		}
	}
	return result
}

func evaluate(testValue int, numbers []int, currentValue int) bool {
	if len(numbers) == 0 {
		return testValue == currentValue
	}
	if evaluate(testValue, numbers[1:], currentValue+numbers[0]) {
		return true
	}
	if evaluate(testValue, numbers[1:], currentValue*numbers[0]) {
		return true
	}
	return false
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
}
