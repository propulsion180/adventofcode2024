package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func is_valid(rules map[int][]int, ruleInv map[int][]int, update []int) bool {
	fmt.Println(update)
	for i, val := range update {
		before := update[:i]
		after := update[i+1:]

		beforeRule := ruleInv[val]
		afterRule := rules[val]

		for _, b := range before {
			if !slices.Contains(beforeRule, b) {
				return false
			}
		}

		for _, a := range after {
			if !slices.Contains(afterRule, a) {
				return false
			}
		}
	}
	return true
}

func MiddleElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	middleIndex := len(nums) / 2
	return nums[middleIndex]
}

func fix_it(rules map[int][]int, rulesInv map[int][]int, wrong []int) int {
	temp := wrong

	for !is_valid(rules, rulesInv, temp) {
		for i, val := range temp {
			if i != len(temp)-1 {
				if slices.Contains(rulesInv[val], temp[i+1]) {
					t := temp[i+1]
					tt := temp[i]

					temp[i] = t
					temp[i+1] = tt
				}
			}
		}

	}

	return MiddleElement(temp)
}

func main() {
	rules := make(map[int][]int)
	rulesInv := make(map[int][]int)
	var updates [][]int

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	collectRules := true
	for scanner.Scan() {
		line := scanner.Text()

		if collectRules {
			if line == "" {
				collectRules = false
				continue
			}
			parts := strings.Split(line, "|")
			if len(parts) < 2 {
				log.Fatal("Error in gettign rules")
			}
			one, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal("Couldn't convert first part to int", err)
			}

			two, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Couldn't convert second part to int", err)
			}

			rulesInv[two] = append(rulesInv[two], one) //so that when you find something you can see what was meant to be before it
			rules[one] = append(rules[one], two)
		} else {
			var temp []int
			stringInts := strings.Split(line, ",")
			for _, val := range stringInts {
				int, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal("Couldn't convert update string to int", err)
				}
				temp = append(temp, int)
			}
			updates = append(updates, temp)
		}
	}
	fmt.Println("Collected")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalMiddles := 0

	fmt.Println("Checking for validity")

	for _, val := range updates {
		if !is_valid(rules, rulesInv, val) {
			fmt.Println(val)
			totalMiddles = totalMiddles + fix_it(rules, rulesInv, val)
		}
	}

	fmt.Println(totalMiddles)
}
