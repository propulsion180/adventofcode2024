package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkX(slice [][]string, x int, y int) bool {
	if x-1 < 0 {
		return false
	}
	if y-1 < 0 {
		return false
	}
	if x+1 > len(slice[y])-1 {
		return false
	}
	if y+1 > len(slice)-1 {
		return false
	}

	upLeft := slice[y-1][x-1]
	upRight := slice[y-1][x+1]
	downLeft := slice[y+1][x-1]
	downRight := slice[y+1][x+1]
	if upLeft == "X" || upRight == "X" || downLeft == "X" || downRight == "X" {
		return false
	}

	if upLeft == "A" || upRight == "A" || downLeft == "A" || downRight == "A" {
		return false
	}

	if upRight == downLeft {
		return false
	}

	if upLeft == downRight {
		return false
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(upLeft, upRight)
	fmt.Println(downLeft, downRight)

	return true
}

func findAllXMASES(slice [][]string) int {
	total := 0
	fmt.Println("HEre")
	for i, val := range slice {
		for j, val2 := range val {
			if val2 == "A" {
				if checkX(slice, j, i) {
					total++
				}
			}
		}
	}
	return total
}

func main() {
	var dataSlice [][]string

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, field := range fields {
			field2 := strings.Split(field, "")
			dataSlice = append(dataSlice, field2)
		}
	}

	fmt.Println(findAllXMASES(dataSlice))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
