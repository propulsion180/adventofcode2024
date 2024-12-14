package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Robot struct {
	x, y   int
	vx, vy int
}

func main() {
	width := 101
	height := 103
	steps := 100
	midX := 50
	midY := 51

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

	for i := 0; i < steps; i++ {
		for r := range robots {
			robots[r].x += robots[r].vx
			robots[r].y += robots[r].vy
			if robots[r].x < 0 {
				robots[r].x = (robots[r].x % width) + width
				if robots[r].x == width {
					robots[r].x = 0
				}
			} else if robots[r].x >= width {
				robots[r].x = robots[r].x % width
			}
			if robots[r].y < 0 {
				robots[r].y = (robots[r].y % height) + height
				if robots[r].y == height {
					robots[r].y = 0
				}
			} else if robots[r].y >= height {
				robots[r].y = robots[r].y % height
			}
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		if r.x < midX && r.y < midY {
			q1++
		} else if r.x > midX && r.y < midY {
			q2++
		} else if r.x < midX && r.y > midY {
			q3++
		} else if r.x > midX && r.y > midY {
			q4++
		}
	}

	fmt.Println(q1 * q2 * q3 * q4)
}
