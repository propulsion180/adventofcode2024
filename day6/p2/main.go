package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func nextPosGen(currX int, currY int, dir string) (int, int) {
	switch dir {
	case "up":
		return currX, currY - 1
	case "down":
		return currX, currY + 1
	case "left":
		return currX - 1, currY
	case "right":
		return currX + 1, currY
	default:
		log.Fatal("Failed to get next direction")
	}
	return 0, 0
}

func turnRight(dir string) string {
	switch dir {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	default:
		log.Fatal("Failed to get the next direction")
	}
	return ""
}

func traverseWithObstruction(input [][]string, startX, startY int, dir string, obsX, obsY int) bool {
	inputCopy := make([][]string, len(input))
	for i := range input {
		inputCopy[i] = append([]string{}, input[i]...)
	}
	inputCopy[obsY][obsX] = "#"

	visited := make(map[string]string)
	currX, currY := startX, startY
	currDir := dir

	for {
		if currY < 0 || currX < 0 || currY >= len(inputCopy) || currX >= len(inputCopy[currY]) {
			break
		}

		key := strconv.Itoa(currX) + "," + strconv.Itoa(currY)
		if prevDir, ok := visited[key]; ok && prevDir == currDir {
			// Detected a loop
			return true
		}
		visited[key] = currDir

		nextX, nextY := nextPosGen(currX, currY, currDir)
		for nextX >= 0 && nextY >= 0 && nextY < len(inputCopy) && nextX < len(inputCopy[nextY]) && inputCopy[nextY][nextX] == "#" {
			currDir = turnRight(currDir)
			nextX, nextY = nextPosGen(currX, currY, currDir)
		}

		currX, currY = nextX, nextY
	}
	return false
}

func findLoopPositions(input [][]string, startX, startY int, dir string) int {
	validPositions := 0
	for y, row := range input {
		for x, cell := range row {
			if cell != "." && cell != "#" {
				continue
			}
			if x == startX && y == startY {
				continue
			}
			if traverseWithObstruction(input, startX, startY, dir, x, y) {
				validPositions++
			}
		}
	}
	return validPositions
}

func main() {
	var input [][]string

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "")
		input = append(input, splitted)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var startX, startY int
	var dir string

	for i, val := range input {
		for j, val2 := range val {
			switch val2 {
			case "^":
				startX, startY, dir = j, i, "up"
			case ">":
				startX, startY, dir = j, i, "right"
			case "<":
				startX, startY, dir = j, i, "left"
			case "v":
				startX, startY, dir = j, i, "down"
			}
		}
	}

	loopPositions := findLoopPositions(input, startX, startY, dir)
	fmt.Println("Number of positions to cause a loop:", loopPositions)
}
