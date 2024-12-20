package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	row, col int
}

type DistanceMap map[Point]int

func main() {
	grid, start, end := readGrid("data.txt")
	result := solve(grid, start, end)
	fmt.Printf("Number of cheats saving at least 100 picoseconds: %d\n", result)
}

func solve(grid [][]rune, start, end Point) int {
	distFromStart := computeDistances(grid, start)
	distToEnd := computeDistances(grid, end)

	normalDist := distFromStart[end]

	count := 0
	rows, cols := len(grid), len(grid[0])
	seen := make(map[string]bool)

	for r1 := 0; r1 < rows; r1++ {
		for c1 := 0; c1 < cols; c1++ {
			if grid[r1][c1] == '#' {
				continue
			}
			pos1 := Point{r1, c1}
			distToPos1 := distFromStart[pos1]
			if distToPos1 == -1 {
				continue
			}

			for r2 := max(0, r1-2); r2 <= min(rows-1, r1+2); r2++ {
				for c2 := max(0, c1-2); c2 <= min(cols-1, c1+2); c2++ {
					if abs(r2-r1)+abs(c2-c1) > 2 {
						continue
					}

					pos2 := Point{r2, c2}
					if grid[r2][c2] == '#' {
						continue
					}

					distFromPos2 := distToEnd[pos2]
					if distFromPos2 == -1 {
						continue
					}

					cheatDist := abs(r2-r1) + abs(c2-c1)
					totalDist := distToPos1 + cheatDist + distFromPos2

					savings := normalDist - totalDist
					if savings >= 100 {
						key := fmt.Sprintf("%d,%d-%d,%d", r1, c1, r2, c2)
						if !seen[key] {
							count++
							seen[key] = true
						}
					}
				}
			}
		}
	}

	return count
}

func computeDistances(grid [][]rune, start Point) DistanceMap {
	distances := make(DistanceMap)
	queue := []Point{start}
	distances[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := Point{current.row + dir.row, current.col + dir.col}

			if !isValid(grid, next) {
				continue
			}

			if _, exists := distances[next]; exists {
				continue
			}

			distances[next] = distances[current] + 1
			queue = append(queue, next)
		}
	}

	return distances
}

func readGrid(filename string) ([][]rune, Point, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	var start, end Point
	scanner := bufio.NewScanner(file)
	row := 0

	for scanner.Scan() {
		line := []rune(scanner.Text())
		for col, ch := range line {
			if ch == 'S' {
				start = Point{row, col}
				line[col] = '.'
			} else if ch == 'E' {
				end = Point{row, col}
				line[col] = '.'
			}
		}
		grid = append(grid, line)
		row++
	}

	return grid, start, end
}

func isValid(grid [][]rune, p Point) bool {
	return p.row >= 0 && p.row < len(grid) &&
		p.col >= 0 && p.col < len(grid[0]) &&
		grid[p.row][p.col] != '#'
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
