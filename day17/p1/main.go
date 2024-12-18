package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")

	regA := 0
	regB := 0
	regC := 0
	var ops []int

	if err != nil {
		log.Fatal("Failed to open file", err)
	}

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
				break
			case "Register B":
				regB, err = strconv.Atoi(strings.TrimSpace(splitByColon[1]))
				if err != nil {
					log.Fatal("Failed to convert string to int")
				}
				break
			case "Register C":
				regC, err = strconv.Atoi(strings.TrimSpace(splitByColon[1]))
				if err != nil {
					log.Fatal("Failed to convert string to int")
				}
				break
			default:
				log.Fatal("Got an extra line that is invalid")
				break
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
		log.Fatal("Error during Scanning", err)
	}

	fmt.Println("RegA: ", strconv.Itoa(regA))
	fmt.Println("RegB: ", strconv.Itoa(regB))
	fmt.Println("RegC: ", strconv.Itoa(regC))
	fmt.Println(ops)

	fmt.Println(runOpsOnRegs(regA, regB, regC, ops))

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
			break
		case 1:
			regB = regB ^ ops[currPos+1]
			currPos += 2
			break
		case 2:
			regB = comboOp(regA, regB, regC, ops[currPos+1]) % 8
			currPos += 2
			break
		case 3:
			if regA == 0 {
				currPos += 2
				break
			} else {
				currPos = ops[currPos+1]
				break
			}
		case 4:
			regB = regB ^ regC
			currPos += 2
			break
		case 5:
			out += strconv.Itoa(comboOp(regA, regB, regC, ops[currPos+1])%8) + ","
			currPos += 2
			break
		case 6:
			bottom := math.Pow(2, float64(comboOp(regA, regB, regC, ops[currPos+1])))
			currPos += 2
			regB = int(float64(regA) / bottom)
			break
		case 7:
			bottom := math.Pow(2, float64(comboOp(regA, regB, regC, ops[currPos+1])))
			currPos += 2
			regC = int(float64(regA) / bottom)
			break
		default:
			log.Fatal("invalid opcode")
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
		log.Fatal("Invalid combo")
	default:
		return op
	}
	return -1
}
