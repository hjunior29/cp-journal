package main

import (
	"fmt"
)

func solve(w int) string {
	if w < 1 || w > 100 {
		return "NO"
	}

	if w%2 == 0 {
		if w > 2 {
			return "YES"
		} else {
			return "NO"
		}
	} else {
		return "NO"
	}
}

func main() {
	var w int
	fmt.Scan(&w)
	result := solve(w)
	fmt.Println(result)
}

// Problem 4A: Watermelon
// Time: 5m
