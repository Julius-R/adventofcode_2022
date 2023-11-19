package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/2

func main() {
	file, err := os.ReadFile("./plays.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	var part1, part2 int
	match := map[string]struct {
		r1, r2 int
	}{
		"A Y": {8, 4},
		"B Z": {9, 9},
		"C X": {7, 2},
		"A Z": {3, 8},
		"B X": {1, 1},
		"C Y": {2, 6},
		"A X": {4, 3},
		"B Y": {5, 5},
		"C Z": {6, 7},
	}

	for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		if line != "" {
			p := match[line]
			part1 += p.r1
			part2 += p.r2
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
