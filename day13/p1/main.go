package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	X, Y   int
}

func main() {
	machines, err := readMachinesFromFile("data.txt")
	if err != nil {
		fmt.Println("Error reading machines:", err)
		return
	}

	maxPresses := 100
	solvedCount := 0
	totalCost := 0

	for _, m := range machines {
		cost := solveMachine(m, maxPresses)
		if cost >= 0 {
			solvedCount++
			totalCost += cost
		}
	}

	fmt.Printf("Maximum number of prizes: %d\n", solvedCount)
	fmt.Printf("Minimum total cost to achieve that: %d\n", totalCost)
}

func solveMachine(m Machine, maxPresses int) int {
	minCost := -1
	for a := 0; a <= maxPresses; a++ {
		for b := 0; b <= maxPresses; b++ {
			xVal := a*m.Ax + b*m.Bx
			yVal := a*m.Ay + b*m.By
			if xVal == m.X && yVal == m.Y {
				cost := 3*a + b
				if minCost == -1 || cost < minCost {
					minCost = cost
				}
			}
		}
	}
	return minCost
}

func readMachinesFromFile(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(lines) == 3 {
				m, err := parseMachine(lines)
				if err != nil {
					return nil, err
				}
				machines = append(machines, m)
				lines = []string{}
			}
			continue
		}
		lines = append(lines, line)
	}

	if len(lines) == 3 {
		m, err := parseMachine(lines)
		if err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	return machines, scanner.Err()
}

func parseMachine(lines []string) (Machine, error) {
	var m Machine

	ax, ay, err := parseButtonLine(lines[0], "A")
	if err != nil {
		return m, err
	}
	bx, by, err := parseButtonLine(lines[1], "B")
	if err != nil {
		return m, err
	}
	px, py, err := parsePrizeLine(lines[2])
	if err != nil {
		return m, err
	}

	m = Machine{
		Ax: ax,
		Ay: ay,
		Bx: bx,
		By: by,
		X:  px,
		Y:  py,
	}

	return m, nil
}

func parseButtonLine(line, buttonType string) (int, int, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("invalid button line format: %s", line)
	}

	afterColon := strings.TrimSpace(parts[1])
	chunks := strings.Split(afterColon, ",")
	if len(chunks) != 2 {
		return 0, 0, fmt.Errorf("invalid button line format: %s", line)
	}

	xPart := strings.TrimSpace(chunks[0])
	yPart := strings.TrimSpace(chunks[1])

	ax, err := parseCoordPart(xPart, 'X')
	if err != nil {
		return 0, 0, err
	}
	ay, err := parseCoordPart(yPart, 'Y')
	if err != nil {
		return 0, 0, err
	}

	return ax, ay, nil
}

func parsePrizeLine(line string) (int, int, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("invalid prize line format: %s", line)
	}

	afterColon := strings.TrimSpace(parts[1])
	chunks := strings.Split(afterColon, ",")
	if len(chunks) != 2 {
		return 0, 0, fmt.Errorf("invalid prize line format: %s", line)
	}

	xPart := strings.TrimSpace(chunks[0])
	yPart := strings.TrimSpace(chunks[1])

	px, err := parsePrizeCoordPart(xPart, 'X')
	if err != nil {
		return 0, 0, err
	}
	py, err := parsePrizeCoordPart(yPart, 'Y')
	if err != nil {
		return 0, 0, err
	}

	return px, py, nil
}

func parseCoordPart(s string, coord rune) (int, error) {
	if len(s) < 2 || s[0] != byte(coord) {
		return 0, fmt.Errorf("invalid coordinate part (expected %c): %s", coord, s)
	}
	valStr := s[1:]
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("invalid integer in coord part: %v", err)
	}
	return val, nil
}

func parsePrizeCoordPart(s string, coord rune) (int, error) {
	if len(s) < 2 || s[0] != byte(coord) || s[1] != '=' {
		return 0, fmt.Errorf("invalid prize coordinate part (expected %c=): %s", coord, s)
	}
	valStr := s[2:]
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("invalid integer in prize coord part: %v", err)
	}
	return val, nil
}
