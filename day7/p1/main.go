package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func generateCombinations(length int) [][]rune {
	if length <= 0 {
		return nil
	}

	var combinations [][]rune
	total := 1 << length

	for i := 0; i < total; i++ {
		var combination []rune
		for j := 0; j < length; j++ {
			if i&(1<<j) != 0 {
				combination = append(combination, '*')
			} else {
				combination = append(combination, '+')
			}
		}
		combinations = append(combinations, combination)
	}

	return combinations
}

func compute(nums []*big.Int, comb []rune) *big.Int {
	result := new(big.Int).Set(nums[0])

	for i := 1; i < len(nums); i++ {
		if comb[i-1] == '+' {
			result.Add(result, nums[i])
		} else {
			result.Mul(result, nums[i])
		}
	}
	return result
}

func processEquations(data map[*big.Int][]*big.Int) *big.Int {
	total := big.NewInt(0)

	for target, numbers := range data {
		operators := generateCombinations(len(numbers) - 1)
		for _, comb := range operators {
			if compute(numbers, comb).Cmp(target) == 0 {
				total.Add(total, target)
				break
			}
		}
	}

	return total
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data := make(map[*big.Int][]*big.Int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid format in line: %s", line)
		}

		target := new(big.Int)
		if _, ok := target.SetString(strings.TrimSpace(parts[0]), 10); !ok {
			log.Fatalf("Failed to parse target value: %s", parts[0])
		}

		numStrings := strings.Fields(parts[1])
		var numbers []*big.Int
		for _, numStr := range numStrings {
			num := new(big.Int)
			if _, ok := num.SetString(numStr, 10); !ok {
				log.Fatalf("Failed to parse number: %s", numStr)
			}
			numbers = append(numbers, num)
		}

		data[target] = numbers
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	result := processEquations(data)
	fmt.Printf("Total Calibration Result: %s\n", result.String())
}
