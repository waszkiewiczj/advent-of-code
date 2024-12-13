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

type Operator func(x, y int) int

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func concat(x, y int) int {
	s := string(x) + string(y)
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func evaluate(testValue int, numbers []int, currentValue int, operators []Operator) bool {
	if currentValue > testValue {
		return false
	}
	if len(numbers) == 0 {
		return testValue == currentValue
	}
	for _, op := range operators {
		if evaluate(testValue, numbers, op(currentValue, numbers[0]), operators) {
			return true
		}
	}
	return false
}

func part1(input []Calibration) int {
	result := 0
	for _, c := range input {
		if evaluate(c.testValue, c.numbers, 0, []Operator{add, mul}) {
			result += c.testValue
		}
	}
	return result
}

func part2(input []Calibration) int {
	result := 0
	for _, c := range input {
		if evaluate(c.testValue, c.numbers, 0, []Operator{add, mul, concat}) {
			result += c.testValue
		}
	}
	return result
}
func main() {
	input := parseInput("test.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
	fmt.Printf("PART 2: %v\n", part2(input))
}
