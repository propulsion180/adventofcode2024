package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, e := os.Open("data.txt")
	if e != nil {
		log.Fatal("Failed to open file")
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var rawData [][][]string

	var temp [][]string
	for s.Scan() {
		l := s.Text()
		if l == "" {
			fmt.Println("is empty")

			rawData = append(rawData, temp)
			temp = nil
			continue
		}
		temp = append(temp, strings.Split(l, ""))
		fmt.Println("added 2d")
	}
	rawData = append(rawData, temp)
	fmt.Println(len(rawData))
	if e := s.Err(); e != nil {
		log.Fatal("Failed to read file")
	}

	locks, keys := splitToLocksAndKeys(rawData)
	fmt.Println(locks)
	fmt.Println(keys)

	totalPlausible := 0

	for _, l := range locks {
		for _, k := range keys {
			if notOverlap(l, k) {
				totalPlausible++
			}
		}
	}

	fmt.Println(totalPlausible)

}

func splitToLocksAndKeys(in [][][]string) ([][]int, [][]int) {
	var keys [][]int
	var locks [][]int

	for _, v := range in {
		if isLock(v) {
			locks = append(locks, lockToInts(v))
		} else {
			keys = append(keys, KeyToInts(v))
		}
	}

	return locks, keys
}

func isLock(in [][]string) bool {
	for _, v := range in[0] {
		if v != "#" {
			return false
		}
	}
	return true
}

func lockToInts(in [][]string) []int {
	temp := []int{0, 0, 0, 0, 0}
	coltotal := 0
	for i := 0; i < len(in[0]); i++ {
		for j := 1; j < len(in); j++ {
			if in[j][i] == "#" {
				coltotal++
			}
		}
		temp[i] = coltotal
		coltotal = 0
	}
	return temp
}

func KeyToInts(in [][]string) []int {
	temp := []int{0, 0, 0, 0, 0}
	coltotal := 0
	for i := 0; i < len(in[0]); i++ {
		for j := 0; j < len(in)-1; j++ {
			if in[j][i] == "#" {
				coltotal++
			}
		}
		temp[i] = coltotal
		coltotal = 0
	}
	return temp
}

func notOverlap(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}
