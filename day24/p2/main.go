package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var gates []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			gates = append(gates, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return gates
}

func findGate(a, b, op string, gates []string) string {
	prefix1 := fmt.Sprintf("%s %s %s", a, op, b)
	prefix2 := fmt.Sprintf("%s %s %s", b, op, a)

	for _, g := range gates {
		if strings.HasPrefix(g, prefix1) || strings.HasPrefix(g, prefix2) {
			parts := strings.Split(g, "->")
			if len(parts) != 2 {
				continue
			}
			outWire := strings.TrimSpace(parts[1])
			return outWire
		}
	}
	return ""
}

func reconstructAdder(gates []string, bitCount int) string {
	var swapped []string

	var carryIn string

	for i := 0; i < bitCount; i++ {
		n := fmt.Sprintf("%02d", i)
		m1 := findGate("x"+n, "y"+n, "XOR", gates)
		n1 := findGate("x"+n, "y"+n, "AND", gates)

		var r1, z1, carryOut string

		if carryIn != "" {
			r1 = findGate(carryIn, m1, "AND", gates)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = findGate(carryIn, m1, "AND", gates)
			}

			z1 = findGate(carryIn, m1, "XOR", gates)

			if strings.HasPrefix(m1, "z") {
				mOld := m1
				m1 = z1
				z1 = mOld
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				nOld := n1
				n1 = z1
				z1 = nOld
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				rOld := r1
				r1 = z1
				z1 = rOld
				swapped = append(swapped, r1, z1)
			}

			carryOut = findGate(r1, n1, "OR", gates)

			if strings.HasPrefix(carryOut, "z") && carryOut != fmt.Sprintf("z%02d", bitCount) {
				coOld := carryOut
				carryOut = z1
				z1 = coOld
				swapped = append(swapped, carryOut, z1)
			}

		} else {
			carryOut = n1
		}

		if carryIn == "" {
			carryIn = n1
		} else {
			carryIn = carryOut
		}
	}

	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}

func main() {
	filename := "data.txt"
	gates := parseInput(filename)

	const bitCount = 45
	result := reconstructAdder(gates, bitCount)

	fmt.Println(result)
}
