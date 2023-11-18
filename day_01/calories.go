package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://adventofcode.com/2022/day/1

func GetCalorieCount() int {
	file, err := os.Open("./cals.txt")
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("error closing file: %s", err)
		}
	}()

	var hc1, hc2, hc3, currentCount int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			switch {
			case currentCount > hc1:
				hc3 = hc2
				hc2 = hc1
				hc1 = currentCount
			case currentCount > hc2:
				hc3 = hc2
				hc2 = currentCount
			case currentCount > hc3:
				hc3 = currentCount
			}
			currentCount = 0
		} else if val, intErr := strconv.Atoi(scanner.Text()); intErr != nil {
			fmt.Printf("error parsing line: %s", intErr)
		} else {
			currentCount += val
		}
	}

	// Handle the last count if the file doesn't end with an empty line
	switch {
	case currentCount > hc1:
		hc3 = hc2
		hc2 = hc1
		hc1 = currentCount
	case currentCount > hc2:
		hc3 = hc2
		hc2 = currentCount
	case currentCount > hc3:
		hc3 = currentCount
	}

	return hc1 + hc2 + hc3
}

func main() {
	fmt.Println(GetCalorieCount())
}
