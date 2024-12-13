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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stonesString := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read from file:", err)
	}

	stones := make(map[int]int)
	for _, val := range strings.Fields(stonesString) {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		stones[num]++
	}

	fmt.Println("Initial stones:", stones)

	stones = runBlinks(stones, 75)

	totalStones := 0
	for _, count := range stones {
		totalStones += count
	}

	fmt.Println("Total stones after 25 blinks:", totalStones)
}

func runBlinks(stones map[int]int, blinks int) map[int]int {
	for i := 0; i < blinks; i++ {
		fmt.Printf("Blink %d\n", i+1)
		stones = blink(stones)
	}
	return stones
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
		} else if lenOfNum(stone)%2 == 0 {
			left, right := splitNum(stone)
			newStones[left] += count
			newStones[right] += count
		} else {
			newStones[stone*2024] += count
		}
	}

	return newStones
}

func lenOfNum(num int) int {
	if num == 0 {
		return 1
	}
	return len(strconv.Itoa(num))
}

func splitNum(num int) (int, int) {
	strNum := strconv.Itoa(num)
	half := len(strNum) / 2

	left, err1 := strconv.Atoi(strNum[:half])
	right, err2 := strconv.Atoi(strNum[half:])
	if err1 != nil || err2 != nil {
		log.Fatal("Failed to split number:", err1, err2)
	}

	return left, right
}
