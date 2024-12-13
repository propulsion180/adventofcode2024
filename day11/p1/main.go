package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stonesString := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read from read file", err)
	}

	var stones []int
	for _, val := range strings.Fields(stonesString) {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		stones = append(stones, num)
	}

	fmt.Println(stones)

	fmt.Println(len(runBlinks(stones)))
}

func runBlinks(in []int) []int {
	temp := make([]int, len(in))
	copy(temp, in)

	for i := 0; i < 25; i++ {
		temp = blink(temp)
	}

	return temp
}

func blink(in []int) []int {
	var temp []int

	for _, val := range in {
		if val == 0 {
			temp = append(temp, 1)
			continue
		}

		if lenOfNum(val)%2 == 0 {
			left, right := splitNum(val)
			temp = append(temp, left)
			temp = append(temp, right)
			continue
		}

		temp = append(temp, val*2024)
	}

	return temp
}

func lenOfNum(in int) int {
	return len(strconv.Itoa(in))
}

func splitNum(in int) (int, int) {
	strNum := strconv.Itoa(in)
	len := len(strNum)
	half := len / 2

	left, err1 := strconv.Atoi(strNum[:half])
	right, err2 := strconv.Atoi(strNum[half:])

	if err1 != nil || err2 != nil {
		log.Fatal("Failed to reconvert halves")
	}

	return left, right
}
