package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Robot struct {
	x, y   int
	vx, vy int
}

func main() {
	width := 101
	height := 103

	delay := 10 * time.Millisecond

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Failed to open data file: %v", err)
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		pPart := strings.TrimPrefix(parts[0], "p=")
		vPart := strings.TrimPrefix(parts[1], "v=")
		var x, y, vx, vy int
		fmt.Sscanf(pPart, "%d,%d", &x, &y)
		fmt.Sscanf(vPart, "%d,%d", &vx, &vy)
		robots = append(robots, Robot{x: x, y: y, vx: vx, vy: vy})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	second := 0
	for {
		grid := make([][]int, height)
		for i := range grid {
			grid[i] = make([]int, width)
		}

		for _, r := range robots {
			grid[r.y][r.x]++
		}

		fmt.Printf("Second: %d\n", second)
		for _, row := range grid {
			for _, val := range row {
				if val == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(val)
				}
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(delay)

		for i := range robots {
			robots[i].x += robots[i].vx
			robots[i].y += robots[i].vy
			if robots[i].x < 0 {
				robots[i].x = (robots[i].x % width) + width
				if robots[i].x == width {
					robots[i].x = 0
				}
			} else if robots[i].x >= width {
				robots[i].x = robots[i].x % width
			}
			if robots[i].y < 0 {
				robots[i].y = (robots[i].y % height) + height
				if robots[i].y == height {
					robots[i].y = 0
				}
			} else if robots[i].y >= height {
				robots[i].y = robots[i].y % height
			}
		}

		second++
	}
}
