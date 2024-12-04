package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := "data.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %v/n", err)
		return
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	re := regexp.MustCompile(`mul\((\-?\d+),(\-?\d+)\)`)

	matches := re.FindAllStringSubmatch(content, -1)

	sum := 0

	for _, match := range matches {
		if len(match) == 3 {
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				product := x * y
				sum += product
				fmt.Printf("Found: mul(%d,%d) -> Product: %d\n", x, y, product)
			} else {
				fmt.Println("Error converting numbers:", match[1], match[2])
			}
		}
	}

	fmt.Println(sum)
}
