package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindOverlap(str string) bool {
	var a []string
	arr := strings.Split(str, ",")
	a = append(a, strings.Split(arr[0], "-")...)
	a = append(a, strings.Split(arr[1], "-")...)

	l1, _ := strconv.Atoi(a[0])
	l2, _ := strconv.Atoi(a[1])
	r1, _ := strconv.Atoi(a[2])
	r2, _ := strconv.Atoi(a[3])

	// p1
	// return l1 <= r1 && l2 >= r2 || l1 >= r1 && l2 <= r2
	// p2
	return l1 <= r2 && l2 >= r1

	// switch {
	// case l1 <= r1 && l2 >= r2:
	// 	fallthrough
	// case l1 >= r1 && l2 <= r2:
	// 	fallthrough
	// case l1 >= r1 && r1 <= l2:
	// 	fallthrough
	// case l1 <= r1 && r1 <= l2:
	// 	fallthrough
	// case l1 <= r2 && r1 >= l2:
	// 	return true
	// default:
	// 	return false
	// }
}

func main() {
	file, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	score := 0
	for _, v := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		if v != "" && FindOverlap(v) {
			score++
		}
	}

	fmt.Println(score)
}
