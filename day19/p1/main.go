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

	fmt.Println(numPossible(towelsAvailable, combosWanted))
}

func numPossible(tA []string, cW []string) int {
	ret := 0

	for _, design := range cW {
		memo := make(map[string]bool)
		if canForm(design, tA, memo) {
			ret++
		}
	}

	return ret
}

func canForm(design string, tA []string, memo map[string]bool) bool {
	if design == "" {
		return true
	}

	if val, found := memo[design]; found {
		return val
	}

	for _, towel := range tA {
		if strings.HasPrefix(design, towel) {
			if canForm(design[len(towel):], tA, memo) {
				memo[design] = true
				return true
			}
		}
	}

	memo[design] = false
	return false
}
