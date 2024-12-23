package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) map[string][]string {
	connections := make(map[string][]string)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		connections[a] = append(connections[a], b)
		connections[b] = append(connections[b], a)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return connections
}

func findGroups(connections map[string][]string) [][]string {
	groups := [][]string{}

	for a, neighborsA := range connections {
		for _, b := range neighborsA {
			if b <= a {
				continue
			}
			for _, c := range connections[b] {
				if c <= b || c == a {
					continue
				}
				if isConnected(connections, c, a) {
					group := []string{a, b, c}
					groups = append(groups, group)
				}
			}
		}
	}

	return groups
}

func isConnected(connections map[string][]string, node1, node2 string) bool {
	for _, neighbor := range connections[node1] {
		if neighbor == node2 {
			return true
		}
	}
	return false
}

func main() {
	filename := "data.txt"
	connections := readInput(filename)
	groups := findGroups(connections)

	count := 0
	for _, group := range groups {
		containsT := false
		for _, computer := range group {
			if strings.HasPrefix(computer, "t") {
				containsT = true
				break
			}
		}
		if containsT {
			fmt.Println(strings.Join(group, ","))
			count++
		}
	}

	fmt.Printf("Total groups containing at least one 't': %d\n", count)
}
