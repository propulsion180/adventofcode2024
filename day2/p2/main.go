package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if diff > 0 {
			decreasing = false
		}
		if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing
}

func canBeSafeWithDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		modified := make([]int, 0, len(levels)-1)
		modified = append(modified, levels[:i]...)
		modified = append(modified, levels[i+1:]...)
		if isSafeReport(modified) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numSafe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		levels := make([]int, len(fields))
		for i, val := range fields {
			currVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("invalid number: %s", val)
			}
			levels[i] = currVal
		}

		if isSafeReport(levels) || canBeSafeWithDampener(levels) {
			numSafe++
		}
	}

	fmt.Println(numSafe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
