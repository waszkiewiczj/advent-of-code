package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinates struct {
	X, Y int
}

type Map struct {
	data          [][]byte
	width, height int
}

func NewMap(data [][]byte) Map {
	height := len(data)
	width := len(data[0])
	return Map{
		data:   data,
		width:  width,
		height: height,
	}
}

func (m *Map) GetAntennas() map[byte][]Coordinates {
	result := make(map[byte][]Coordinates)
	for x, row := range m.data {
		for y, antenna := range row {
			if antenna == '.' {
				continue
			}
			antennas, ok := result[antenna]
			if !ok {
				antennas = []Coordinates{}
			}
			result[antenna] = append(antennas, Coordinates{x, y})
		}
	}
	return result
}

func parseInput(filepath string) Map {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	data := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []byte(line))
	}
	return NewMap(data)
}

func part1(m Map) int {
	antinodes := [][]bool{}
	for i := 0; i < m.height; i++ {
		antinodes = append(antinodes, make([]bool, m.width))
	}
	result := 0
	for _, coords := range m.GetAntennas() {
		for idx, c1 := range coords[:len(coords)-1] {
			for _, c2 := range coords[idx+1:] {
				if c1.X > c2.X {
					c1, c2 = c2, c1
				}
				distX := c2.X - c1.X
				distY := c2.Y - c1.Y

				antenna1 := Coordinates{c1.X - distX, c1.Y - distY}
				antenna2 := Coordinates{c2.X + distX, c2.Y + distY}
				for _, antenna := range []Coordinates{antenna1, antenna2} {
					if antenna.X >= 0 && antenna.X < m.height && antenna.Y >= 0 && antenna.Y < m.width {
						if !antinodes[antenna.X][antenna.Y] {
							result++
						}
						antinodes[antenna.X][antenna.Y] = true
					}
				}
			}
		}
	}
	return result
}

func part2(m Map) int {
	antinodes := [][]int{}
	for i := 0; i < m.height; i++ {
		antinodes = append(antinodes, make([]int, m.width))
	}
	result := 0
	for _, coords := range m.GetAntennas() {
		result++
		for idx, c1 := range coords[:len(coords)-1] {
			result++
			for _, c2 := range coords[idx+1:] {
				if c1.X > c2.X {
					c1, c2 = c2, c1
				}
				distX := c2.X - c1.X
				distY := c2.Y - c1.Y

				antenna1 := Coordinates{c1.X - distX, c1.Y - distY}
				for antenna1.X >= 0 && antenna1.X < m.height && antenna1.Y >= 0 && antenna1.Y < m.width {
					if antinodes[antenna1.X][antenna1.Y] == 0 && m.data[antenna1.X][antenna1.Y] == '.' {
						result++
					}
					antinodes[antenna1.X][antenna1.Y]++
					antenna1.X -= distX
					antenna1.Y -= distY
				}

				antenna2 := Coordinates{c2.X + distX, c2.Y + distY}
				for antenna2.X >= 0 && antenna2.X < m.height && antenna2.Y >= 0 && antenna2.Y < m.width {
					if antinodes[antenna2.X][antenna2.Y] == 0 && m.data[antenna2.X][antenna2.Y] == '.' {
						result++
					}
					antinodes[antenna2.X][antenna2.Y]++
					antenna2.X += distX
					antenna2.Y += distY
				}
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
