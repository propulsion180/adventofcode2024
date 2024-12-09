package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func parseInput(input []string) map[Point]rune {
	antennas := make(map[Point]rune)
	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antennas[Point{X: x, Y: y}] = char
			}
		}
	}
	return antennas
}

func calculateAntinodes(antennas map[Point]rune, width, height int) map[Point]struct{} {
	antinodes := make(map[Point]struct{})
	for p1, f1 := range antennas {
		for p2, f2 := range antennas {
			if p1 == p2 || f1 != f2 {
				continue
			}
			dx := p2.X - p1.X
			dy := p2.Y - p1.Y
			antinode1 := Point{X: p2.X + dx, Y: p2.Y + dy}
			antinode2 := Point{X: p1.X - dx, Y: p1.Y - dy}
			if antinode1.X >= 0 && antinode1.X < width && antinode1.Y >= 0 && antinode1.Y < height {
				antinodes[antinode1] = struct{}{}
			}
			if antinode2.X >= 0 && antinode2.X < width && antinode2.Y >= 0 && antinode2.Y < height {
				antinodes[antinode2] = struct{}{}
			}
		}
	}
	return antinodes
}

func countAntinodes(input []string) int {
	width := len(input[0])
	height := len(input)
	antennas := parseInput(input)
	antinodes := calculateAntinodes(antennas, width, height)
	return len(antinodes)
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			input = append(input, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := countAntinodes(input)
	fmt.Println("Number of unique antinodes:", result)
}
