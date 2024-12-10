package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	topographicMap, err := readMapFromFile("data.txt")
	if err != nil {
		fmt.Println("Error reading the map:", err)
		return
	}

	sumOfScores := computeTrailheadScores(topographicMap)

	fmt.Println("Sum of trailhead scores:", sumOfScores)
}

func readMapFromFile(filename string) ([][]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	mapData := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, ch := range line {
			row[j], _ = strconv.Atoi(string(ch))
		}
		mapData[i] = row
	}

	return mapData, nil
}

func computeTrailheadScores(mapData [][]int) int {
	totalScore := 0

	for r := range mapData {
		for c := range mapData[r] {
			if mapData[r][c] == 0 {
				totalScore += countReachableNines(mapData, r, c)
			}
		}
	}

	return totalScore
}

func countReachableNines(mapData [][]int, startR, startC int) int {
	rows, cols := len(mapData), len(mapData[0])
	stack := []struct{ r, c int }{{startR, startC}}
	visited := make(map[string]bool)
	reachableNines := make(map[string]bool)

	key := func(r, c int) string { return fmt.Sprintf("%d,%d", r, c) }

	visited[key(startR, startC)] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if mapData[current.r][current.c] == 9 {
			reachableNines[key(current.r, current.c)] = true
			continue
		}

		for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nr, nc := current.r+dir[0], current.c+dir[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[key(nr, nc)] {
				if mapData[nr][nc] == mapData[current.r][current.c]+1 {
					visited[key(nr, nc)] = true
					stack = append(stack, struct{ r, c int }{nr, nc})
				}
			}
		}
	}

	return len(reachableNines)
}
