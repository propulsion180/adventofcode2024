package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func findLargestClique(connections map[string][]string) []string {
	largestClique := []string{}

	for node := range connections {
		candidateClique := []string{node}
		neighbors := connections[node]

		for _, neighbor := range neighbors {
			if isFullyConnected(connections, candidateClique, neighbor) {
				candidateClique = append(candidateClique, neighbor)
			}
		}

		if len(candidateClique) > len(largestClique) {
			largestClique = candidateClique
		}
	}

	sort.Strings(largestClique)
	return largestClique
}

func isFullyConnected(connections map[string][]string, clique []string, node string) bool {
	for _, member := range clique {
		if !isConnected(connections, member, node) {
			return false
		}
	}
	return true
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
	largestClique := findLargestClique(connections)

	if len(largestClique) == 0 {
		fmt.Println("No LAN party clique found.")
	} else {
		password := strings.Join(largestClique, ",")
		fmt.Printf("Password to get into the LAN party: %s\n", password)
	}
}
