package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// File reading and parsing
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	defer file.Close()

	regA := 0
	regB := 0
	regC := 0
	var ops []int

	scanner := bufio.NewScanner(file)
	operations := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			operations = true
			continue
		}
		if !operations {
			splitByColon := strings.Split(line, ":")
			reg := splitByColon[0]
			switch reg {
			case "Register A":
				regA, err = strconv.Atoi(strings.TrimSpace(splitByColon[1]))
				if err != nil {
					log.Fatal("Failed to convert string to int")
				}
			case "Register B":
				regB, err = strconv.Atoi(strings.TrimSpace(splitByColon[1]))
				if err != nil {
					log.Fatal("Failed to convert string to int")
				}
			case "Register C":
				regC, err = strconv.Atoi(strings.TrimSpace(splitByColon[1]))
				if err != nil {
					log.Fatal("Failed to convert string to int")
				}
			default:
				log.Fatal("Got an extra line that is invalid")
			}
		} else {
			getOps := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ",")
			for _, op := range getOps {
				temp, err := strconv.Atoi(op)
				if err != nil {
					log.Fatal("Failed to convert opcode to int")
				}
				ops = append(ops, temp)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error during scanning", err)
	}

	fmt.Println(regA)

	// Target output
	targetOutput := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ops)), ","), "[]") + ","

	// Parallelized trials
	const maxA = 1000000000000000 // Upper limit for search (for demo purposes)
	const workers = 12            // Number of goroutines to use
	chunkSize := maxA / workers

	result := make(chan int, 1) // Channel to receive the answer
	var wg sync.WaitGroup       // WaitGroup to manage goroutines

	// Split the search range into chunks for each worker
	for i := 0; i < workers; i++ {
		start := i * chunkSize
		eind := start + chunkSize
		if i == workers-1 {
			end = maxA // Ensure the last worker covers the remaining range
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for a := start; a < end; a++ {
				// Print progress every 1,000,000 iterations
				if a%1000000 == 0 {
					fmt.Printf("Worker [%d-%d]: Trying Register A = %d\n", start, end, a)
				}

				curr := runOpsOnRegs(a, regB, regC, ops)
				if curr == targetOutput {
					result <- a // Send the answer to the channel
					return      // Exit the goroutine
				}
			}
		}(start, end)
	}

	// Wait for all workers in a separate goroutine
	go func() {
		wg.Wait()
		close(result)
	}()

	// Read the result
	if val, ok := <-result; ok {
		fmt.Println("Found the correct value for Register A:", val)
	} else {
		fmt.Println("No valid value for Register A was found.")
	}
}

func runOpsOnRegs(a, b, c int, ops []int) string {
	var regA int = a
	var regB int = b
	var regC int = c

	currPos := 0
	out := ""

	for currPos < len(ops) {
		op := ops[currPos]
		switch op {
		case 0:
			bottom := math.Pow(2, float64(comboOp(regA, regB, regC, ops[currPos+1])))
			currPos += 2
			regA = int(float64(regA) / bottom)
		case 1:
			regB = regB ^ ops[currPos+1]
			currPos += 2
		case 2:
			regB = comboOp(regA, regB, regC, ops[currPos+1]) % 8
			currPos += 2
		case 3:
			if regA == 0 {
				currPos += 2
			} else {
				currPos = ops[currPos+1]
			}
		case 4:
			regB = regB ^ regC
			currPos += 2
		case 5:
			out += strconv.Itoa(comboOp(regA, regB, regC, ops[currPos+1])%8) + ","
			currPos += 2
		case 6:
			bottom := math.Pow(2, float64(comboOp(regA, regB, regC, ops[currPos+1])))
			currPos += 2
			regB = int(float64(regA) / bottom)
		case 7:
			bottom := math.Pow(2, float64(comboOp(regA, regB, regC, ops[currPos+1])))
			currPos += 2
			regC = int(float64(regA) / bottom)
		default:
			log.Fatal("Invalid opcode")
		}
	}
	return out
}

func comboOp(a, b, c, op int) int {
	switch op {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	case 7:
		log.Fatal("Invalid combo operand")
	default:
		return op
	}
	return -1
}

