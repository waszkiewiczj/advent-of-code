package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parseInput(filepath string) ([][]int, error) {
	result := make([][]int, 2)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for idx, numberRaw := range strings.Fields(line) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				return nil, err
			}
			result[idx] = append(result[idx], number)
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func part1(input [][]int) int {
	left := append(IntHeap{}, input[0]...)
	heap.Init(&left)
	right := append(IntHeap{}, input[1]...)
	heap.Init(&right)

	result := 0
	for left.Len() > 0 && right.Len() > 0 {
		l := heap.Pop(&left).(int)
		r := heap.Pop(&right).(int)
		diff := int(math.Abs(float64(l) - float64(r)))
		result += diff
	}
	return result
}

func part2(input [][]int) int {
	counts := make(map[int]int)
	for _, r := range input[1] {
		counts[r]++
	}

	result := 0
	for _, l := range input[0] {
		count, ok := counts[l]
		if !ok {
			count = 0
		}
		result += (count * l)
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
