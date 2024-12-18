package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const gridSize = 71
const maxBytes = 1024

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	memory := make([][]rune, gridSize)
	for i := range memory {
		memory[i] = make([]rune, gridSize)
		for j := range memory[i] {
			memory[i][j] = '.'
		}
	}

	scanner := bufio.NewScanner(file)
	bytePositions := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		bytePositions = append(bytePositions, Point{x, y})
	}

	for i := 0; i < maxBytes && i < len(bytePositions); i++ {
		p := bytePositions[i]
		memory[p.y][p.x] = '#'
	}

	start := Point{0, 0}
	end := Point{gridSize - 1, gridSize - 1}
	steps := findShortestPath(memory, start, end)

	if steps != -1 {
		fmt.Printf("The minimum number of steps needed is: %d\n", steps)
	} else {
		fmt.Println("No path to the exit exists.")
	}
}

func findShortestPath(memory [][]rune, start, end Point) int {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make(map[Point]bool)
	queue := []struct {
		point Point
		steps int
	}{{start, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.point == end {
			return curr.steps
		}

		if visited[curr.point] {
			continue
		}
		visited[curr.point] = true

		for _, dir := range directions {
			neighbor := Point{curr.point.x + dir.x, curr.point.y + dir.y}
			if isValid(neighbor, memory, visited) {
				queue = append(queue, struct {
					point Point
					steps int
				}{neighbor, curr.steps + 1})
			}
		}
	}

	return -1
}

func isValid(p Point, memory [][]rune, visited map[Point]bool) bool {
	return p.x >= 0 && p.x < gridSize && p.y >= 0 && p.y < gridSize &&
		memory[p.y][p.x] != '#' && !visited[p]
}
