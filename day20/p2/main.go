package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
)

type Point struct {
	row, col int
}

type Distances map[Point]int

func main() {
	grid, start, end := readGrid("data.txt")
	result := solve(grid, start, end)
	fmt.Printf("Number of cheats saving at least 100 picoseconds: %d\n", result)
}

func solve(grid [][]rune, start, end Point) int {
	distFromStart := computeDistances(grid, start)
	distToEnd := computeDistances(grid, end)
	normalDist := distFromStart[end]

	rows, cols := len(grid), len(grid[0])
	var count int64 = 0

	numWorkers := runtime.NumCPU()
	jobs := make(chan Point, rows*cols)
	var wg sync.WaitGroup
	seenLock := &sync.Mutex{}
	seen := make(map[string]bool)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for startPos := range jobs {
				localCount := processPoint(grid, startPos, distFromStart, distToEnd, normalDist, seen, seenLock)
				atomic.AddInt64(&count, localCount)
			}
		}()
	}

	// Send jobs
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '#' {
				pos := Point{r, c}
				if dist := distFromStart[pos]; dist != -1 {
					jobs <- pos
				}
			}
		}
	}
	close(jobs)

	wg.Wait()
	return int(count)
}

func processPoint(grid [][]rune, startPos Point, distFromStart, distToEnd Distances, normalDist int, seen map[string]bool, seenLock *sync.Mutex) int64 {
	var localCount int64
	maxSteps := 20
	distToStart := distFromStart[startPos]

	type QueueItem struct {
		pos   Point
		steps int
	}

	queue := []QueueItem{{startPos, 0}}
	visited := make(map[Point]bool)
	visited[startPos] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.steps > 0 && grid[current.pos.row][current.pos.col] != '#' {
			distFromEnd := distToEnd[current.pos]
			if distFromEnd != -1 {
				totalDist := distToStart + current.steps + distFromEnd
				savings := normalDist - totalDist

				if savings >= 100 {
					key := fmt.Sprintf("%d,%d-%d,%d", startPos.row, startPos.col, current.pos.row, current.pos.col)
					seenLock.Lock()
					if !seen[key] {
						seen[key] = true
						localCount++
					}
					seenLock.Unlock()
				}
			}
		}

		if current.steps >= maxSteps {
			continue
		}

		// Try all four directions
		for _, dir := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := Point{current.pos.row + dir.row, current.pos.col + dir.col}

			if !isInBounds(grid, next) || visited[next] {
				continue
			}

			visited[next] = true
			queue = append(queue, QueueItem{next, current.steps + 1})
		}
	}

	return localCount
}

func computeDistances(grid [][]rune, start Point) Distances {
	distances := make(Distances)
	queue := []Point{start}
	distances[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := Point{current.row + dir.row, current.col + dir.col}

			if !isInBounds(grid, next) || grid[next.row][next.col] == '#' {
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

func isInBounds(grid [][]rune, p Point) bool {
	return p.row >= 0 && p.row < len(grid) &&
		p.col >= 0 && p.col < len(grid[0])
}
