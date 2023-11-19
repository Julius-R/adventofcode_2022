package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/3

func GetItemType(str string) int {
	split := len([]rune(str)) / 2
	seenCharacters := map[rune]bool{}
	for _, v := range []rune(str[:split]) {
		if !seenCharacters[v] {
			seenCharacters[v] = true
		}
	}
	for _, v := range []rune(str[split:]) {
		if seenCharacters[v] {
			return GetValue(string(v))
		}
	}
	return 0
}

func GetGroupItemType(str []string) int {
	var triple string
	seenCharacters := map[rune]bool{}
	for _, v := range []rune(str[0]) {
		if !seenCharacters[v] {
			seenCharacters[v] = true
		}
	}
	for _, v := range []rune(str[1]) {
		if seenCharacters[v] {
			triple += string(v)
		}
	}

	for _, v := range []rune(str[2]) {
		if strings.Contains(triple, string(v)) {
			return GetValue(string(v))
		}
	}
	return 0
}

func GetValue(str string) int {
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

	var points, points2, count int
	var args []string

	for _, str := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		if str != "" {
			points += GetItemType(str)
			args = append(args, str)
			if len(args) == 3 {
				count = 0
				points2 += GetGroupItemType(args)
				args = []string{}
			}
			count++
		}
	}

	fmt.Println(points2)
}
