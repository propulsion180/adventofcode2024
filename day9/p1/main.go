package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func expand(in []int) []int {
	var temp []int
	value := true
	index := 0
	for _, val := range in {
		if value {
			value = false
			for j := 0; j < val; j++ {
				temp = append(temp, index)
			}
			index++
		} else {
			value = true
			for j := 0; j < val; j++ {
				temp = append(temp, -1)
			}
		}
	}

	return temp
}

func defragment(in []int) []int {
	temp := in

	l := 0
	r := len(temp) - 1

	for l < r {
		lVal := in[l]
		rVal := in[r]

		if lVal != -1 {
			l++
			continue
		}

		if rVal == -1 {
			r--
			continue
		}

		in[r] = lVal
		in[l] = rVal
		l++
		r--
	}
	return temp
}

func checksum(in []int) int {
	tot := 0

	for i, val := range in {
		if val < 0 {
			continue
		}

		tot = tot + (i * val)
	}

	return tot
}

func main() {
	input := ""
	var inputInt []int
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input += line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	splitInput := strings.Split(input, "")

	for _, val := range splitInput {
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Failed to convert to int")
		}
		inputInt = append(inputInt, num)
	}

	fmt.Println(inputInt)

	expanded := expand(inputInt)

	fmt.Println(expanded)

	defraged := defragment(expanded)

	fmt.Println(defraged)

	sum := checksum(defraged)

	fmt.Println(sum)

}
