package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filepath string) ([][]int, [][]int) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules := [][]int{}
	updates := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineRaw := scanner.Text()
		if strings.Contains(lineRaw, "|") {
			rule := []int{}
			for _, numberRaw := range strings.Split(lineRaw, "|") {
				number, err := strconv.Atoi(numberRaw)
				if err != nil {
					panic(err)
				}
				rule = append(rule, number)
			}
			rules = append(rules, rule)
		} else if lineRaw != "" {
			update := []int{}
			for _, numberRaw := range strings.Split(lineRaw, ",") {
				number, err := strconv.Atoi(numberRaw)
				if err != nil {
					panic(err)
				}
				update = append(update, number)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func isEqualRule(r1, r2 []int) bool {
	for i := 0; i < min(len(r1), len(r2)); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return len(r1) == len(r2)
}

func containsRule(rule []int, rules [][]int) bool {
	for _, r := range rules {
		if isEqualRule(rule, r) {
			return true
		}
	}
	return false
}

func isValidUpdate(rules [][]int, update []int) bool {
	for idx, n1 := range update {
		for _, n2 := range update[idx+1:] {
			rule := []int{n1, n2}
			if !containsRule(rule, rules) {
				return false
			}
		}
	}
	return true
}

func part1(rules, updates [][]int) int {
	result := 0

	for _, update := range updates {
		if isValidUpdate(rules, update) {
			result += update[len(update)/2]
		}
	}

	return result
}

type SortableUpdate struct {
	update []int
	rules  [][]int
}

func (s SortableUpdate) Len() int { return len(s.update) }

func (s SortableUpdate) Swap(i, j int) { s.update[i], s.update[j] = s.update[j], s.update[i] }

func (s SortableUpdate) Less(i, j int) bool {
	for _, rule := range s.rules {
		if s.update[i] == rule[0] && s.update[j] == rule[1] {
			return true
		}
	}
	return false
}

func fixUpdate(rules [][]int, update []int) []int {
	s := SortableUpdate{
		update: update,
		rules:  rules,
	}
	sort.Sort(s)
	return s.update
}

func part2(rules, updates [][]int) int {
	result := 0

	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			validUpdate := fixUpdate(rules, update)
			result += validUpdate[len(validUpdate)/2]
		}
	}

	return result
}

func main() {
	inputRules, inputUpdates := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(inputRules, inputUpdates))
	fmt.Printf("PART 2: %v\n", part2(inputRules, inputUpdates))
}
