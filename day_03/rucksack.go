package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/3

func GetItemType(str string) int {
	split := len([]rune(str)) / 2
	duplicateItem := map[rune]bool{}
	for _, v := range []rune(str[:split]) {
		if duplicateItem[v] {
			continue
		}
		duplicateItem[v] = true
	}
	for _, v := range []rune(str[split:]) {
		if duplicateItem[v] {
			return GetPrioritySum(string(v))
		}
	}
	return 0
}

func GetPrioritySum(str string) int {
	count := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 12,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	var total int
	if comparison := strings.Compare(str, strings.ToUpper(str)); comparison == 0 {
		total = count[strings.ToLower(str)] + 26
	} else {
		total = count[str]
	}
	return total
}

func main() {
	file, err := os.ReadFile("./rucks.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	var points int

	for _, str := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		if str != "" {
			points += GetItemType(str)
		}
	}

	fmt.Println(points)
}
