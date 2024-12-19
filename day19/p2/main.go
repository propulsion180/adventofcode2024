package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	defer f.Close()

	var towelsAvailable []string
	var combosWanted []string

	sec1 := true

	s := bufio.NewScanner(f)

	for s.Scan() {
		l := s.Text()

		if sec1 {
			if l == "" {
				sec1 = false
				continue
			}
			towelsAvailable = strings.Split(l, ", ")
		} else {
			combosWanted = append(combosWanted, l)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal("Failed to read")
	}

	fmt.Println(totalCombinations(towelsAvailable, combosWanted))
}

func totalCombinations(tA []string, cW []string) int {
	total := 0

	for _, design := range cW {
		memo := make(map[string]int)
		total += countWays(design, tA, memo)
	}

	return total
}

func countWays(design string, tA []string, memo map[string]int) int {
	if design == "" {
		return 1
	}

	if val, found := memo[design]; found {
		return val
	}

	ways := 0

	for _, towel := range tA {
		if strings.HasPrefix(design, towel) {
			ways += countWays(design[len(towel):], tA, memo)
		}
	}

	memo[design] = ways
	return ways
}
