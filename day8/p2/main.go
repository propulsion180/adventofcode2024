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

			current := p2
			for {
				current = Point{X: current.X + dx, Y: current.Y + dy}
				if current.X < 0 || current.X >= width || current.Y < 0 || current.Y >= height {
					break
				}
				antinodes[current] = struct{}{}
			}

			current = p1
			for {
				current = Point{X: current.X - dx, Y: current.Y - dy}
				if current.X < 0 || current.X >= width || current.Y < 0 || current.Y >= height {
					break
				}
				antinodes[current] = struct{}{}
			}
		}
	}

	for p := range antennas {
		antinodes[p] = struct{}{}
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
