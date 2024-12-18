package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const gridSize = 71

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

	start := Point{0, 0}
	end := Point{gridSize - 1, gridSize - 1}

	for _, p := range bytePositions {
		memory[p.y][p.x] = '#'

		if !pathExists(memory, start, end) {
			fmt.Printf("%d,%d\n", p.x, p.y)
			return
		}
	}
}

func pathExists(memory [][]rune, start, end Point) bool {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make(map[Point]bool)
	queue := []Point{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return true
		}

		if visited[curr] {
			continue
		}
		visited[curr] = true

		for _, dir := range directions {
			neighbor := Point{curr.x + dir.x, curr.y + dir.y}
			if isValid(neighbor, memory, visited) {
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

func isValid(p Point, memory [][]rune, visited map[Point]bool) bool {
	return p.x >= 0 && p.x < gridSize && p.y >= 0 && p.y < gridSize &&
		memory[p.y][p.x] != '#' && !visited[p]
}
