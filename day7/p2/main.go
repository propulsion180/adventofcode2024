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
	total := intPow(3, length)

	for i := 0; i < total; i++ {
		var combination []rune
		val := i
		for j := 0; j < length; j++ {
			mod := val % 3
			if mod == 0 {
				combination = append(combination, '+')
			} else if mod == 1 {
				combination = append(combination, '*')
			} else {
				combination = append(combination, '|')
			}
			val /= 3
		}
		combinations = append(combinations, combination)
	}

	return combinations
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func compute(nums []*big.Int, comb []rune) *big.Int {
	result := new(big.Int).Set(nums[0])
	concatBuffer := ""

	for i := 1; i < len(nums); i++ {
		op := comb[i-1]
		if op == '+' {
			result.Add(result, nums[i])
		} else if op == '*' {
			result.Mul(result, nums[i])
		} else if op == '|' {
			if concatBuffer == "" {
				concatBuffer = result.String()
			}
			concatBuffer += nums[i].String()
			concatValue := new(big.Int)
			concatValue.SetString(concatBuffer, 10)
			result.Set(concatValue)
			concatBuffer = ""
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
