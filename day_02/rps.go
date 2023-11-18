package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/2

func GetRoundScore(str string) int {
	match := map[string]string{
		"A Y": "WIN",
		"B Z": "WIN",
		"C X": "WIN",
		"A Z": "LOSE",
		"B X": "LOSE",
		"C Y": "LOSE",
		"A X": "DRAW",
		"B Y": "DRAW",
		"C Z": "DRAW",
	}

	roundDecision := map[string]int{
		"LOSE": 0,
		"DRAW": 3,
		"WIN":  6,
	}

	choicePoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	userChoice := string(str[2])
	return roundDecision[match[str]] + choicePoints[userChoice]
}

func GetTotalScore() int {
	file, err := os.ReadFile("./plays.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	var totalScore int
	for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		if line != "" {
			totalScore += GetRoundScore(line)
		}
	}
	return totalScore
}

func main() {
	fmt.Println(GetTotalScore())
}
