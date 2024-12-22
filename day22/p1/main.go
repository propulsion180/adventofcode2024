package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer f.Close()

	var secNums []int

	s := bufio.NewScanner(f)

	for s.Scan() {
		l := s.Text()
		lInt, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal("Failed to convert secret number to int")
		}
		secNums = append(secNums, lInt)
	}

	if err := s.Err(); err != nil {
		log.Fatal("error in reading file")
	}

	total := 0

	for _, secret := range secNums {
		var toAdd int = secret
		for i := 0; i < 2000; i++ {
			toAdd = step3(step2(step1(toAdd)))
		}
		total += toAdd
		fmt.Printf("Secret No: %d, final: %d \n", secret, toAdd)
	}

	fmt.Printf("Total: %d \n", total)

}

func step1(sec int) int {
	preMix := sec * 64
	mixed := mix(preMix, sec)
	return prune(mixed)
}

func step2(sec int) int {
	div := float64(sec / 32)
	floored := math.Floor(div)
	mixed := mix(int(floored), sec)
	return prune(mixed)
}

func step3(sec int) int {
	mult := sec * 2048
	mixed := mix(mult, sec)
	return prune(mixed)
}

func prune(in int) int {
	return in % 16777216
}

func mix(in, sec int) int {
	return in ^ sec
}
