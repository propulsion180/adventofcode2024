package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	defer file.Close()

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
	fmt.Println("Areas:", areas)

	sides := findSides(data, regions)
	fmt.Println("Sides:", sides)

	fmt.Println("Total price:", getPrices(areas, sides))
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

func findSides(data [][]string, regions map[string][]pos) map[string]int {
	sides := make(map[string]int)

	for regionID, cells := range regions {
		shape := make(map[pos]bool)
		for _, cell := range cells {
			shape[cell] = true
		}

		count := 0
		rows := len(data)
		cols := len(data[0])

		for col := 0; col < cols; col++ {
			for row := 0; row < rows; row++ {
				if shape[pos{col, row}] && !shape[pos{col - 1, row}] {
					count++
					row++
					for ; row < rows; row++ {
						if !(shape[pos{col, row}] && !shape[pos{col - 1, row}]) {
							break
						}
					}
				}
			}
		}

		for col := cols - 1; col >= 0; col-- {
			for row := 0; row < rows; row++ {
				if shape[pos{col, row}] && !shape[pos{col + 1, row}] {
					count++
					row++
					for ; row < rows; row++ {
						if !(shape[pos{col, row}] && !shape[pos{col + 1, row}]) {
							break
						}
					}
				}
			}
		}

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				if shape[pos{col, row}] && !shape[pos{col, row - 1}] {
					count++
					col++
					for ; col < cols; col++ {
						if !(shape[pos{col, row}] && !shape[pos{col, row - 1}]) {
							break
						}
					}
				}
			}
		}

		for row := rows - 1; row >= 0; row-- {
			for col := 0; col < cols; col++ {
				if shape[pos{col, row}] && !shape[pos{col, row + 1}] {
					count++
					col++
					for ; col < cols; col++ {
						if !(shape[pos{col, row}] && !shape[pos{col, row + 1}]) {
							break
						}
					}
				}
			}
		}

		sides[regionID] = count
	}

	return sides
}

func getPrices(area map[string]int, sides map[string]int) int {
	total := 0
	for k, area := range area {
		total += area * sides[k]
	}
	return total
}
