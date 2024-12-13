package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Ax, Ay int64
	Bx, By int64
	X, Y   int64
}

func main() {
	machines, err := readMachinesFromFile("data.txt")
	if err != nil {
		fmt.Println("Error reading machines:", err)
		return
	}

	const offset = 10000000000000
	for i := range machines {
		machines[i].X += offset
		machines[i].Y += offset
	}

	solvedCount := 0
	totalCost := int64(0)

	for _, m := range machines {
		cost := solveMachine(m)
		if cost >= 0 {
			solvedCount++
			totalCost += cost
		}
	}

	fmt.Printf("Maximum number of prizes: %d\n", solvedCount)
	fmt.Printf("Minimum total cost to achieve that: %d\n", totalCost)
}

func solveMachine(m Machine) int64 {
	D := m.Ax*m.By - m.Bx*m.Ay

	if D == 0 {
		return -1
	}

	numA := m.X*m.By - m.Y*m.Bx
	numB := m.Ax*m.Y - m.Ay*m.X

	if numA%D != 0 || numB%D != 0 {
		return -1
	}

	a := numA / D
	b := numB / D

	if a < 0 || b < 0 {
		return -1
	}

	cost := 3*a + b
	return cost
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

func parseButtonLine(line, buttonType string) (int64, int64, error) {
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

func parsePrizeLine(line string) (int64, int64, error) {
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

func parseCoordPart(s string, coord rune) (int64, error) {
	if len(s) < 2 || rune(s[0]) != coord {
		return 0, fmt.Errorf("invalid coordinate part (expected %c): %s", coord, s)
	}
	valStr := s[1:]
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid integer in coord part: %v", err)
	}
	return val, nil
}

func parsePrizeCoordPart(s string, coord rune) (int64, error) {
	if len(s) < 3 || rune(s[0]) != coord || s[1] != '=' {
		return 0, fmt.Errorf("invalid prize coordinate part (expected %c=): %s", coord, s)
	}
	valStr := s[2:]
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid integer in prize coord part: %v", err)
	}
	return val, nil
}
