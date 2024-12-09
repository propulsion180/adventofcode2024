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

func defragmentWhole(in []int) []int {
	temp := make([]int, len(in))
	copy(temp, in)
	maxFileId := 0
	for _, val := range temp {
		if val >= 0 && val > maxFileId {
			maxFileId = val
		}
	}
	for fileId := maxFileId; fileId >= 0; fileId-- {
		var fileBlocks []int
		for i, val := range temp {
			if val == fileId {
				fileBlocks = append(fileBlocks, i)
			}
		}
		if len(fileBlocks) == 0 {
			continue
		}
		fileStart := fileBlocks[0]
		freeStart := -1
		for i := 0; i < fileStart; i++ {
			if temp[i] == -1 {
				currentFree := 1
				for j := i + 1; j < fileStart && temp[j] == -1; j++ {
					currentFree++
				}
				if currentFree >= len(fileBlocks) {
					freeStart = i
					break
				}
			}
		}
		if freeStart != -1 {
			for _, block := range fileBlocks {
				temp[block] = -1
			}
			for i := 0; i < len(fileBlocks); i++ {
				temp[freeStart+i] = fileId
			}
		}
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
	expanded := expand(inputInt)
	defraged := defragmentWhole(expanded)
	sum := checksum(defraged)
	fmt.Println("Checksum:", sum)
}
