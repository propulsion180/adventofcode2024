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
	fmt.Println("turning")
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
		log.Fatal("failed to get the next position")
	}

	return ""
}

func traverse(inputi [][]string, startX int, startY int, dir string) int {
	input := inputi
	fmt.Println("Inside traverse")
	var seen int
	currX := startX
	currY := startY
	currDir := dir
	for {
		if currY < 0 {
			break
		}
		if currX < 0 {
			break
		}
		if currY > len(input)-1 {
			break
		}
		if currX > len(input[currY])-1 {
			break
		}
		posCode := strconv.Itoa(currX) + strconv.Itoa(currY)
		fmt.Println(posCode)
		currChar := input[currY][currX]
		if currChar != "X" {
			input[currY][currX] = "X"
			seen++
		}

		fmt.Println(input[currY][currX])
		nextX, nextY := nextPosGen(currX, currY, currDir)
		if nextX >= 0 {
			if nextY >= 0 {
				if nextY < len(input) {
					if nextX < len(input[nextY]) {
						for input[nextY][nextX] == "#" {
							currDir = turnRight(currDir)
							nextX, nextY = nextPosGen(currX, currY, currDir)
						}
					}
				}
			}
		}

		fmt.Println(currDir)
		currX, currY = nextX, nextY
	}
	fmt.Println(input)
	return seen
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

	fmt.Println(input)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var startX, startY int
	var dir string

	for i, val := range input {
		for j, val2 := range val {
			if val2 == "^" {
				startX = j
				startY = i
				dir = "up"
				break
			}
			if val2 == ">" {
				startX = j
				startY = i
				dir = "right"
				break
			}
			if val2 == "<" {
				startX = j
				startY = i
				dir = "left"
				break
			}
			if val2 == "v" {
				startX = j
				startY = i
				dir = "down"
				break
			}
		}
	}

	fmt.Println(startX)
	fmt.Println(startY)

	fmt.Println(traverse(input, startX, startY, dir))
}
