package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 2 {
			log.Fatal("invalid line: %s", line)
		}

		l, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal("invalid num for l: %s", fields[0])
		}

		r, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal("Invalid num for r: %s", fields[1])
		}
		left = append(left, l)
		fmt.Println("size l: ", len(left), "newest: ", l)
		right = append(right, r)
		fmt.Println("size r: ", len(right), "newest: ", r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	if len(left) != len(right) {
		log.Fatal("left and right not same length l: %d, r: %d", len(left), len(right))
	}

	sim_score := 0

	for _, vall := range left {
		times_seen := 0
		for _, valr := range right {
			if valr == vall {
				times_seen++
			}
		}
		sim_score = sim_score + (vall * times_seen)
	}

	fmt.Println(sim_score)
}
