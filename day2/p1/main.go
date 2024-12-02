package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findDiffAndSlope(i, j int) (int, bool) {
	res := i - j
	if res < 0 {
		return res * -1, false
	}
	return res, true
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numSafe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		lineSafe := true
		var prevSlope bool
		for i, val := range fields {
			if i == 0 {
				continue
			}

			currVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid number", val)
			}

			prevVal, err := strconv.Atoi(fields[i-1])
			if err != nil {
				log.Fatal("invalid number", fields[i-1])
			}

			diff, slope := findDiffAndSlope(prevVal, currVal)
			fmt.Println(diff)

			if diff > 3 || diff == 0 {
				lineSafe = false
				break
			}

			if i == 1 {
				prevSlope = slope
				continue
			}

			if slope != prevSlope {
				lineSafe = false
				break
			}
		}
		if lineSafe {
			numSafe++
		}
	}

	fmt.Println(numSafe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
