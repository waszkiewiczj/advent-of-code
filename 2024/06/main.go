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

func (c *Coordinates) clone() Coordinates {
	return Coordinates{
		X: c.X,
		Y: c.Y,
	}
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

func (m *Map) get(c Coordinates) byte {
	return m.data[c.X][c.Y]
}

func (m *Map) set(c Coordinates, b byte) {
	line := []rune(m.data[c.X])
	line[c.Y] = rune(b)
	m.data[c.X] = string(line)
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

var startDirection Coordinates = Coordinates{-1, 0}

type PositionLog struct {
	position, direction Coordinates
}

type Guard struct {
	startPosition, position, direction Coordinates
	log                                []PositionLog
}

func (g *Guard) rotate() {
	g.direction = Coordinates{g.direction.Y, -g.direction.X}
}

func (g *Guard) move(m Map) {
	g.log = append(g.log, PositionLog{
		position:  g.position.clone(),
		direction: g.direction.clone(),
	})
	newPosition := Coordinates{
		g.position.X + g.direction.X,
		g.position.Y + g.direction.Y,
	}
	if m.inBounds(newPosition) {
		if m.get(newPosition) == '#' {
			g.rotate()
			return
		}
		m.visit(newPosition)
	}
	g.position = newPosition
}

func (g *Guard) reset() {
	g.position = g.startPosition.clone()
	g.direction = startDirection.clone()
	g.log = []PositionLog{}
}

func (g *Guard) isLoop() bool {
	for _, l := range g.log {
		if l.direction == g.direction && l.position == g.position {
			return true
		}
	}
	return false
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

func initData(input []string) (Guard, Map) {
	pos := findStartingPosition(input)
	g := Guard{
		startPosition: pos,
		position:      pos,
		direction:     startDirection,
		log:           []PositionLog{},
	}
	m := NewMap(input)
	m.visit(pos)
	return g, m
}

func part1(input []string) int {
	g, m := initData(input)
	for m.inBounds(g.position) {
		g.move(m)
	}
	result := m.countVisited()
	return result
}

func isLoop(g Guard, m Map) bool {
	g.reset()
	for m.inBounds(g.position) {
		g.move(m)
		if g.isLoop() {
			return true
		}
	}
	return false
}

func part2(input []string) int {
	result := 0
	g, m := initData(input)
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			c := Coordinates{i, j}
			if g.startPosition != c {
				oldChar := m.get(c)
				m.set(c, '#')
				if isLoop(g, m) {
					result++
				}
				m.set(c, oldChar)
			}

		}
	}
	return result
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("PART 1: %v\n", part1(input))
	fmt.Printf("PART 2: %v\n", part2(input))
}
