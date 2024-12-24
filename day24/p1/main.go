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
	f, e := os.Open("data.txt")
	if e != nil {
		log.Fatal("Failed to open file", e)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	inputs := true

	wires := make(map[string]string)

	for s.Scan() {
		l := s.Text()
		if inputs {
			if l == "" {
				inputs = false
				continue
			}

			split := strings.Split(l, ": ")
			wires[split[0]] = split[1]
		} else {
			split := strings.Split(l, " -> ")
			wires[split[1]] = split[0]
		}
	}

	if e := s.Err(); e != nil {
		log.Fatal("Failed to read file", e)
	}

	fmt.Println(wires)

	findAllVals(wires)

	fmt.Println(wires)

	sortedZKeys := retSortedZKeys(wires)

	endNum := ""

	for _, val := range sortedZKeys {
		endNum = wires[val] + endNum
		fmt.Printf("%s: %s\n", val, wires[val])
	}

	fmt.Println(endNum)
	fmt.Println(binToDec(endNum))

}

func findAllVals(wires map[string]string) {

	for k, v := range wires {
		if isLeaf(v) {
			continue
		}
		wires[k] = solveBool(v, wires)
	}
}

func solveBool(equation string, wires map[string]string) string {
	one, op, two := splitEqu(equation)

	if !isLeaf(wires[one]) {
		oneIs := solveBool(wires[one], wires)
		wires[one] = oneIs
	}

	if !isLeaf(wires[two]) {
		twoIs := solveBool(wires[two], wires)
		wires[two] = twoIs
	}

	var res string

	switch op {
	case "XOR":
		res = xor(wires[one], wires[two])
		break
	case "OR":
		res = or(wires[one], wires[two])
		break
	case "AND":
		res = and(wires[one], wires[two])
	default:
		log.Fatal("Incorrect Operation")
	}

	return res
}

func splitEqu(equ string) (string, string, string) {
	split := strings.Fields(equ)
	return split[0], split[1], split[2]
}

func xor(one, two string) string {
	if one != two {
		return "1"
	}
	return "0"
}

func and(one, two string) string {
	if one == "1" && two == "1" {
		return "1"
	}
	return "0"
}

func or(one, two string) string {
	if one == "1" || two == "1" {
		return "1"
	}
	return "0"
}

func isLeaf(in string) bool {
	return in == "1" || in == "0"
}

func retSortedZKeys(in map[string]string) []string {
	var toSort []string
	for k, _ := range in {
		if strings.HasPrefix(k, "z") {
			toSort = append(toSort, k)
		}
	}

	sort.Slice(toSort, func(i, j int) bool {
		num1, _ := strconv.Atoi(toSort[i][1:])
		num2, _ := strconv.Atoi(toSort[j][1:])
		return num1 < num2
	})

	return toSort
}

func binToDec(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Fatal("Failed to convert bin to dec")
	}
	return int(dec)
}
