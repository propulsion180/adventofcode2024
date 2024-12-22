package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer f.Close()

	var secNums []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		lInt, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal("Failed to convert secret number to int")
		}
		secNums = append(secNums, lInt)
	}

	if err := s.Err(); err != nil {
		log.Fatal("Error in reading file")
	}

	precomputed := precomputePricesAndChanges(secNums, 2000)

	bestSequence, maxBananas := findBestSequenceOptimized(precomputed)

	fmt.Printf("Best Sequence: %v\n", bestSequence)
	fmt.Printf("Max Bananas: %d\n", maxBananas)
}

type PrecomputedData struct {
	prices  [][]int
	changes [][]int
}

func precomputePricesAndChanges(secNums []int, steps int) PrecomputedData {
	prices := make([][]int, len(secNums))
	changes := make([][]int, len(secNums))

	for i, secret := range secNums {
		prices[i] = generatePrices(secret, steps)
		changes[i] = calculateChanges(prices[i])
	}

	return PrecomputedData{prices: prices, changes: changes}
}

func findBestSequenceOptimized(precomputed PrecomputedData) ([]int, int) {
	var bestSequence []int
	maxBananas := 0

	results := make(chan struct {
		sequence []int
		bananas  int
	}, 1)

	var wg sync.WaitGroup

	sequences := generateSequences()
	batchSize := len(sequences) / 8

	for i := 0; i < len(sequences); i += batchSize {
		end := i + batchSize
		if end > len(sequences) {
			end = len(sequences)
		}
		batch := sequences[i:end]

		wg.Add(1)
		go func(batch [][]int) {
			defer wg.Done()
			for _, seq := range batch {
				totalBananas := calculateBananasForSequence(precomputed, seq)
				results <- struct {
					sequence []int
					bananas  int
				}{seq, totalBananas}
			}
		}(batch)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.bananas > maxBananas {
			maxBananas = result.bananas
			bestSequence = result.sequence
		}
	}

	return bestSequence, maxBananas
}

func calculateBananasForSequence(precomputed PrecomputedData, sequence []int) int {
	totalBananas := 0

	for i := range precomputed.prices {
		price := findFirstMatch(precomputed.prices[i], precomputed.changes[i], sequence)
		totalBananas += price
	}

	return totalBananas
}

func generatePrices(secret, steps int) []int {
	prices := make([]int, steps)
	current := secret

	for i := 0; i < steps; i++ {
		current = step3(step2(step1(current)))
		prices[i] = current % 10
	}

	return prices
}

func calculateChanges(prices []int) []int {
	changes := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		changes[i-1] = prices[i] - prices[i-1]
	}
	return changes
}

func findFirstMatch(prices, changes, sequence []int) int {
	for i := 0; i <= len(changes)-len(sequence); i++ {
		if matchesSequence(changes[i:i+len(sequence)], sequence) {
			return prices[i+len(sequence)]
		}
	}
	return 0
}

func matchesSequence(changes, sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if changes[i] != sequence[i] {
			return false
		}
	}
	return true
}

func generateSequences() [][]int {
	sequences := [][]int{}
	for a := -9; a <= 9; a++ {
		for b := -9; b <= 9; b++ {
			for c := -9; c <= 9; c++ {
				for d := -9; d <= 9; d++ {
					sequences = append(sequences, []int{a, b, c, d})
				}
			}
		}
	}
	return sequences
}

func step1(sec int) int {
	preMix := sec * 64
	mixed := mix(preMix, sec)
	return prune(mixed)
}

func step2(sec int) int {
	div := float64(sec / 32)
	floored := math.Floor(div)
	mixed := mix(int(floored), sec)
	return prune(mixed)
}

func step3(sec int) int {
	mult := sec * 2048
	mixed := mix(mult, sec)
	return prune(mixed)
}

func prune(in int) int {
	return in % 16777216
}

func mix(in, sec int) int {
	return in ^ sec
}
