package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type pos struct {
	x int
	y int
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Failed to open the file")
	}

	scanner := bufio.NewScanner(file)

	var data [][]string

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Split(line, ""))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	regions := findRegions(data)

	areas := findArea(regions)
	fmt.Println(areas)

	perimeters := findPerimeter(data, regions)
	fmt.Println(perimeters)

	fmt.Println(getPrices(areas, perimeters))

}

func findRegions(grid [][]string) map[string][]pos {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	regions := make(map[string][]pos)
	idCounter := 1
	directions := []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var floodFill func(int, int, string) []pos
	floodFill = func(x, y int, plant string) []pos {
		region := []pos{}
		queue := []pos{{x, y}}
		visited[y][x] = true

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			region = append(region, current)

			for _, d := range directions {
				nx, ny := current.x+d.x, current.y+d.y
				if nx >= 0 && ny >= 0 && ny < rows && nx < cols && !visited[ny][nx] && grid[ny][nx] == plant {
					visited[ny][nx] = true
					queue = append(queue, pos{nx, ny})
				}
			}
		}
		return region
	}

	for y, row := range grid {
		for x, plant := range row {
			if !visited[y][x] {
				region := floodFill(x, y, plant)
				regionID := fmt.Sprintf("region_%d", idCounter)
				regions[regionID] = region
				idCounter++
			}
		}
	}

	return regions
}

func findArea(in map[string][]pos) map[string]int {
	temp := make(map[string]int)

	for k, v := range in {
		temp[k] = len(v)
	}

	return temp
}

func findPerimeter(data [][]string, in map[string][]pos) map[string]int {
	temp := make(map[string]int)

	for k, v := range in {
		subTotal := 0
		for _, val := range v {
			subSubTotal := 4
			x, y := val.x, val.y

			up := pos{x: x, y: y - 1}
			left := pos{x: x - 1, y: y}
			right := pos{x: x + 1, y: y}
			down := pos{x: x, y: y + 1}

			if slices.Contains(v, up) {
				subSubTotal--
			}
			if slices.Contains(v, left) {
				subSubTotal--
			}
			if slices.Contains(v, right) {
				subSubTotal--
			}
			if slices.Contains(v, down) {
				subSubTotal--
			}

			subTotal += subSubTotal
		}
		temp[k] = subTotal
	}

	return temp
}

func getPrices(area map[string]int, perimeter map[string]int) int {
	total := 0

	for k, v := range area {
		areaPrice := v * perimeter[k]
		total += areaPrice
	}

	return total
}
