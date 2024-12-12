package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput(filepath string) []string {
	result := []string{}
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

type Coordinates struct {
	X, Y int
}

type Map struct {
	data          []string
	width, height int
	visited       [][]bool
}

func NewMap(data []string) Map {
	m := Map{
		data:    data,
		width:   len(data[0]),
		height:  len(data),
		visited: [][]bool{},
	}
	for _, row := range data {
		m.visited = append(m.visited, make([]bool, len(row)))
	}
	return m
}

func (m *Map) visit(c Coordinates) {
	m.visited[c.X][c.Y] = true
}

func (m *Map) get(c Coordinates) string {
	return string(m.data[c.X][c.Y])
}

func (m *Map) inBounds(c Coordinates) bool {
	return c.X >= 0 && c.X < m.height && c.Y >= 0 && c.Y < m.width
}

func (m *Map) countVisited() int {
	result := 0
	for _, row := range m.visited {
		for _, field := range row {
			if field {
				result++
			}
		}
	}
	return result
}

type Guard struct {
	position, direction Coordinates
}

func (g *Guard) rotate() {
	g.direction = Coordinates{g.direction.Y, -g.direction.X}
}

func (g *Guard) move(m Map) {
	newPosition := Coordinates{
		g.position.X + g.direction.X,
		g.position.Y + g.direction.Y,
	}
	if m.inBounds(newPosition) {
		if m.get(newPosition) == "#" {
			g.rotate()
			return
		}
		m.visit(newPosition)
	}
	g.position = newPosition
}

func findStartingPosition(m []string) Coordinates {
	for i, row := range m {
		for j, field := range row {
			if string(field) == "^" {
				return Coordinates{i, j}
			}
		}
	}
	return Coordinates{-1, -1}
}

func part1(input []string) int {
	pos := findStartingPosition(input)
	g := Guard{position: pos, direction: Coordinates{-1, 0}}
	m := NewMap(input)
	m.visit(pos)
	for m.inBounds(g.position) {
		g.move(m)
	}
	result := m.countVisited()
	return result
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
}
