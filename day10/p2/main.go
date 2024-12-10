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

	sumOfRatings := computeTrailheadRatings(topographicMap)

	fmt.Println("Sum of trailhead ratings:", sumOfRatings)
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

func computeTrailheadRatings(mapData [][]int) int {
	totalRating := 0

	for r := range mapData {
		for c := range mapData[r] {
			if mapData[r][c] == 0 {
				totalRating += countDistinctTrails(mapData, r, c)
			}
		}
	}

	return totalRating
}

func countDistinctTrails(mapData [][]int, startR, startC int) int {
	rows, cols := len(mapData), len(mapData[0])
	trailCount := 0

	var dfs func(r, c, currentHeight int)
	dfs = func(r, c, currentHeight int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || mapData[r][c] != currentHeight {
			return
		}

		if currentHeight == 9 {
			trailCount++
			return
		}

		original := mapData[r][c]
		mapData[r][c] = -1

		for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nr, nc := r+dir[0], c+dir[1]
			dfs(nr, nc, currentHeight+1)
		}

		mapData[r][c] = original
	}

	dfs(startR, startC, 0)
	return trailCount
}
