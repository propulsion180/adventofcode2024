package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkUp(slice [][]string, x int, y int, starting string) bool {
	fmt.Println(starting)
	if y-3 < 0 {
		return false
	}

	if starting == "X" {
		if slice[y-1][x] != "M" {
			return false
		}
		if slice[y-2][x] != "A" {
			return false
		}
		if slice[y-3][x] != "S" {
			return false
		}
	} else {
		if slice[y-1][x] != "A" {
			return false
		}
		if slice[y-2][x] != "M" {
			return false
		}
		if slice[y-3][x] != "X" {
			return false
		}
	}

	return true
}

func checkDown(slice [][]string, x int, y int, starting string) bool {
	if y+3 > len(slice)-1 {
		return false
	}

	if starting == "X" {
		if slice[y+1][x] != "M" {
			return false
		}
		if slice[y+2][x] != "A" {
			return false
		}
		if slice[y+3][x] != "S" {
			return false
		}
	} else {
		if slice[y+1][x] != "A" {
			return false
		}
		if slice[y+2][x] != "M" {
			return false
		}
		if slice[y+3][x] != "X" {
			return false
		}
	}

	return true
}

func checkLeft(slice [][]string, x int, y int, starting string) bool {
	if x-3 < 0 {
		return false
	}

	if starting == "X" {
		if slice[y][x-1] != "M" {
			return false
		}
		if slice[y][x-2] != "A" {
			return false
		}
		if slice[y][x-3] != "S" {
			return false
		}
	} else {
		if slice[y][x-1] != "A" {
			return false
		}
		if slice[y][x-2] != "M" {
			return false
		}
		if slice[y][x-3] != "X" {
			return false
		}
	}

	return true
}

func checkRight(slice [][]string, x int, y int, starting string) bool {
	if x+3 > len(slice[y])-1 {
		return false
	}

	if starting == "X" {
		if slice[y][x+1] != "M" {
			return false
		}
		if slice[y][x+2] != "A" {
			return false
		}
		if slice[y][x+3] != "S" {
			return false
		}
	} else {
		if slice[y][x+1] != "A" {
			return false
		}
		if slice[y][x+2] != "M" {
			return false
		}
		if slice[y][x+3] != "X" {
			return false
		}
	}

	return true
}

func checkUpLeft(slice [][]string, x int, y int, starting string) bool {
	if x-3 < 0 {
		return false
	}

	if y-3 < 0 {
		return false
	}

	if starting == "X" {
		if slice[y-1][x-1] != "M" {
			return false
		}
		if slice[y-2][x-2] != "A" {
			return false
		}
		if slice[y-3][x-3] != "S" {
			return false
		}
	} else {
		if slice[y-1][x-1] != "A" {
			return false
		}
		if slice[y-2][x-2] != "M" {
			return false
		}
		if slice[y-3][x-3] != "X" {
			return false
		}
	}

	return true
}

func checkUpRight(slice [][]string, x int, y int, starting string) bool {
	if x+3 > len(slice[y])-1 {
		return false
	}

	if y-3 < 0 {
		return false
	}

	if starting == "X" {
		if slice[y-1][x+1] != "M" {
			return false
		}
		if slice[y-2][x+2] != "A" {
			return false
		}
		if slice[y-3][x+3] != "S" {
			return false
		}
	} else {
		if slice[y-1][x+1] != "A" {
			return false
		}
		if slice[y-2][x+2] != "M" {
			return false
		}
		if slice[y-3][x+3] != "X" {
			return false
		}
	}

	return true
}

func checkDownLeft(slice [][]string, x int, y int, starting string) bool {
	if x-3 < 0 {
		return false
	}

	if y+3 > len(slice)-1 {
		return false
	}

	if starting == "X" {
		if slice[y+1][x-1] != "M" {
			return false
		}
		if slice[y+2][x-2] != "A" {
			return false
		}
		if slice[y+3][x-3] != "S" {
			return false
		}
	} else {
		if slice[y+1][x-1] != "A" {
			return false
		}
		if slice[y+2][x-2] != "M" {
			return false
		}
		if slice[y+3][x-3] != "X" {
			return false
		}
	}

	return true
}

func checkDownRight(slice [][]string, x int, y int, starting string) bool {
	if x+3 > len(slice[y])-1 {
		return false
	}

	if y+3 > len(slice)-1 {
		return false
	}

	if starting == "X" {
		if slice[y+1][x+1] != "M" {
			return false
		}
		if slice[y+2][x+2] != "A" {
			return false
		}
		if slice[y+3][x+3] != "S" {
			return false
		}
	} else {
		if slice[y+1][x+1] != "A" {
			return false
		}
		if slice[y+2][x+2] != "M" {
			return false
		}
		if slice[y+3][x+3] != "X" {
			return false
		}
	}

	return true
}

func findAllXMASES(slice [][]string) int {
	total := 0
	fmt.Println("HEre")
	for i, val := range slice {
		for j, val2 := range val {
			fmt.Println(val2)
			if val2 == "X" || val2 == "S" {
				if checkUp(slice, j, i, val2) {
					total++
				}

				if checkDown(slice, j, i, val2) {
					total++
				}

				if checkLeft(slice, j, i, val2) {
					total++
				}

				if checkRight(slice, j, i, val2) {
					total++
				}

				if checkUpLeft(slice, j, i, val2) {
					total++
				}

				if checkUpRight(slice, j, i, val2) {
					total++
				}

				if checkDownLeft(slice, j, i, val2) {
					total++
				}

				if checkDownRight(slice, j, i, val2) {
					total++
				}
			}
		}
	}
	return total
}

func main() {
	var dataSlice [][]string

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, field := range fields {
			field2 := strings.Split(field, "")
			dataSlice = append(dataSlice, field2)
		}
	}

	fmt.Println(findAllXMASES(dataSlice) / 2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
