package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Warehouse struct {
	grid       [][]rune
	robotRow   int
	robotCol   int
	numRows    int
	numCols    int
	moveBuffer string
}

var deltas = map[rune][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func loadData(filename string) (*Warehouse, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("empty input file")
	}

	mapWidth := len(lines[0])
	var mapLines []string
	var movesLines []string

lineLoop:
	for i, l := range lines {
		if len(l) != mapWidth {
			movesLines = lines[i:]
			break lineLoop
		}
		for _, c := range l {
			if c != '#' && c != 'O' && c != '@' && c != '.' {
				movesLines = lines[i:]
				break lineLoop
			}
		}
		mapLines = append(mapLines, l)
	}

	if len(mapLines) == len(lines) {
		movesLines = []string{}
	}

	numRows := len(mapLines)
	if numRows == 0 {
		return nil, fmt.Errorf("no map lines found")
	}
	numCols := len(mapLines[0])

	grid := make([][]rune, numRows)
	var robotRow, robotCol int
	robotFound := false
	for r := 0; r < numRows; r++ {
		grid[r] = []rune(mapLines[r])
		for c := 0; c < numCols; c++ {
			if grid[r][c] == '@' {
				robotRow, robotCol = r, c
				robotFound = true
			}
		}
	}
	if !robotFound {
		return nil, fmt.Errorf("no robot '@' found on the map")
	}

	moveBuffer := ""
	for _, line := range movesLines {
		moveBuffer += strings.TrimSpace(line)
	}

	w := &Warehouse{
		grid:       grid,
		robotRow:   robotRow,
		robotCol:   robotCol,
		numRows:    numRows,
		numCols:    numCols,
		moveBuffer: moveBuffer,
	}
	return w, nil
}

func (w *Warehouse) moveRobot(dir rune) {
	drdc := deltas[dir]
	dr, dc := drdc[0], drdc[1]

	newR := w.robotRow + dr
	newC := w.robotCol + dc

	if w.grid[newR][newC] == '#' {
		return
	}

	if w.grid[newR][newC] == '.' {
		w.grid[w.robotRow][w.robotCol] = '.'
		w.grid[newR][newC] = '@'
		w.robotRow, w.robotCol = newR, newC
		return
	}

	if w.grid[newR][newC] == 'O' {
		boxPositions := w.collectBoxes(newR, newC, dir)
		if w.canPushBoxes(boxPositions, dir) {
			w.doPushBoxes(boxPositions, dir)
			w.grid[w.robotRow][w.robotCol] = '.'
			w.grid[newR][newC] = '@'
			w.robotRow, w.robotCol = newR, newC
		} else {
			return
		}
	}
}

func (w *Warehouse) collectBoxes(r, c int, dir rune) [][2]int {
	drdc := deltas[dir]
	dr, dc := drdc[0], drdc[1]

	var boxes [][2]int
	curR, curC := r, c
	for {
		if w.grid[curR][curC] == 'O' {
			boxes = append(boxes, [2]int{curR, curC})
			curR += dr
			curC += dc
		} else {
			break
		}
	}
	return boxes
}

func (w *Warehouse) canPushBoxes(boxes [][2]int, dir rune) bool {
	drdc := deltas[dir]
	dr, dc := drdc[0], drdc[1]

	boxSet := make(map[[2]int]bool)
	for _, b := range boxes {
		boxSet[b] = true
	}

	for i := len(boxes) - 1; i >= 0; i-- {
		br, bc := boxes[i][0], boxes[i][1]
		nr, nc := br+dr, bc+dc
		if w.grid[nr][nc] == '#' {
			return false
		}
		if w.grid[nr][nc] == 'O' {
			// Check if this 'O' is part of the same chain
			if !boxSet[[2]int{nr, nc}] {
				// If it's an O not in the chain, can't push
				return false
			}
		}
	}
	return true
}

func (w *Warehouse) doPushBoxes(boxes [][2]int, dir rune) {
	drdc := deltas[dir]
	dr, dc := drdc[0], drdc[1]

	for i := len(boxes) - 1; i >= 0; i-- {
		br, bc := boxes[i][0], boxes[i][1]
		nr, nc := br+dr, bc+dc
		w.grid[br][bc] = '.'
		w.grid[nr][nc] = 'O'
	}
}

func (w *Warehouse) calcSum() int {
	sum := 0
	for r := 0; r < w.numRows; r++ {
		for c := 0; c < w.numCols; c++ {
			if w.grid[r][c] == 'O' {
				sum += 100*r + c
			}
		}
	}
	return sum
}

func main() {
	w, err := loadData("data.txt")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	for _, move := range w.moveBuffer {
		switch move {
		case '^', 'v', '<', '>':
			w.moveRobot(move)
		default:
		}
	}

	fmt.Println(w.calcSum())
}
