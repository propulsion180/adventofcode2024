package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type keypadType int

const (
	topDirectional keypadType = iota
	middleDirectional
	bottomDirectional
	numeric
)

var directionalButtons = map[[2]int]rune{
	{0, 1}: '^',
	{0, 2}: 'A',
	{1, 0}: '<',
	{1, 1}: 'v',
	{1, 2}: '>',
}
var numericButtons = map[[2]int]rune{
	{0, 0}: '7', {0, 1}: '8', {0, 2}: '9',
	{1, 0}: '4', {1, 1}: '5', {1, 2}: '6',
	{2, 0}: '1', {2, 1}: '2', {2, 2}: '3',
	{3, 1}: '0', {3, 2}: 'A',
}

var startPosTopDir = [2]int{0, 2}
var startPosMidDir = [2]int{0, 2}
var startPosBotDir = [2]int{0, 2}
var startPosNumeric = [2]int{3, 2}

var deltas = map[rune][2]int{
	'^': {-1, 0},
	'v': {+1, 0},
	'<': {0, -1},
	'>': {0, +1},
}

type state struct {
	pos1    [2]int
	pos2    [2]int
	pos3    [2]int
	pos4    [2]int
	codeIdx int
}

type queueElem struct {
	st   state
	dist int
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("cannot open data.txt: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var codes []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		codes = append(codes, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(codes) != 5 {
		log.Fatalf("expected 5 lines in data.txt, got %d", len(codes))
	}

	total := 0
	for _, code := range codes {
		dist := bfsShortestPressesForCode(code)
		numericPart := parseNumericPart(code)
		complexity := dist * numericPart
		total += complexity
	}

	fmt.Println(total)
}

func parseNumericPart(code string) int {
	s := strings.TrimLeft(code[:len(code)-1], "0")
	if s == "" {
		return 0
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot parse numeric part of %q: %v", code, err)
	}
	return val
}

func bfsShortestPressesForCode(code string) int {
	start := state{
		pos1:    startPosTopDir,
		pos2:    startPosMidDir,
		pos3:    startPosBotDir,
		pos4:    startPosNumeric,
		codeIdx: 0,
	}

	visited := make(map[state]bool)
	q := []queueElem{{start, 0}}
	visited[start] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.st.codeIdx == len(code) {
			return curr.dist
		}

		for pos, label := range directionalButtons {

			if !canMoveArmOnDirKeypad(curr.st.pos1, pos) {
				continue
			}

			newPos2, newPos3, newPos4, newCodeIdx, ok := cascadePress(
				curr.st.pos2,
				curr.st.pos3,
				curr.st.pos4,
				curr.st.codeIdx,
				label,
				code,
			)
			if !ok {
				continue
			}

			newSt := state{
				pos1:    pos,
				pos2:    newPos2,
				pos3:    newPos3,
				pos4:    newPos4,
				codeIdx: newCodeIdx,
			}
			if !visited[newSt] {
				visited[newSt] = true
				q = append(q, queueElem{newSt, curr.dist + 1})
			}
		}
	}

	return -1
}

func canMoveArmOnDirKeypad(from, to [2]int) bool {
	if _, ok := directionalButtons[from]; !ok {
		return false
	}
	if _, ok := directionalButtons[to]; !ok {
		return false
	}
	return true
}

func cascadePress(
	pos2, pos3, pos4 [2]int,
	codeIdx int,
	label rune,
	code string,
) ([2]int, [2]int, [2]int, int, bool) {

	newPos2 := pos2
	newPos3 := pos3
	newPos4 := pos4
	newCodeIdx := codeIdx

	if label != 'A' {
		delta := deltas[label]
		candidate := [2]int{pos2[0] + delta[0], pos2[1] + delta[1]}
		if _, ok := directionalButtons[candidate]; !ok {
			return newPos2, newPos3, newPos4, newCodeIdx, false
		}
		newPos2 = candidate
	} else {
		btn2 := directionalButtons[pos2]
		var ok bool
		newPos3, newPos4, newCodeIdx, ok = pressOnDirectionalKeypad(
			pos3, pos4, codeIdx, code,
			btn2,
		)
		if !ok {
			return newPos2, newPos3, newPos4, newCodeIdx, false
		}
	}

	return newPos2, newPos3, newPos4, newCodeIdx, true
}

func pressOnDirectionalKeypad(
	pos3, pos4 [2]int,
	codeIdx int,
	code string,
	label rune,
) ([2]int, [2]int, int, bool) {
	newPos3 := pos3
	newPos4 := pos4
	newCodeIdx := codeIdx
	if label != 'A' {
		// move pos3
		delta := deltas[label]
		candidate := [2]int{pos3[0] + delta[0], pos3[1] + delta[1]}
		if _, ok := directionalButtons[candidate]; !ok {
			return newPos3, newPos4, newCodeIdx, false
		}
		newPos3 = candidate
	} else {
		btn3 := directionalButtons[pos3]
		var ok bool
		newPos4, newCodeIdx, ok = pressOnNumericKeypad(pos4, codeIdx, code, btn3)
		if !ok {
			return newPos3, newPos4, newCodeIdx, false
		}
	}
	return newPos3, newPos4, newCodeIdx, true
}

func pressOnNumericKeypad(
	pos4 [2]int,
	codeIdx int,
	code string,
	label rune,
) ([2]int, int, bool) {
	newPos4 := pos4
	newCodeIdx := codeIdx

	if label != 'A' {
		delta := deltas[label]
		candidate := [2]int{pos4[0] + delta[0], pos4[1] + delta[1]}
		if _, ok := numericButtons[candidate]; !ok {
			return newPos4, newCodeIdx, false
		}
		newPos4 = candidate
	} else {
		btn4 := numericButtons[pos4]
		if codeIdx >= len(code) {
			return newPos4, newCodeIdx, false
		}
		expected := rune(code[codeIdx])
		if btn4 != expected {
			return newPos4, newCodeIdx, false
		}
		newCodeIdx++
	}
	return newPos4, newCodeIdx, true
}
