package main

import (
	"fmt"
	"strings"
)

func main() {
	var instructions int
	var operations []string

	fmt.Scan(&instructions)
	if instructions < 1 || instructions > 150 {
		fmt.Println("Invalid number of instructions")
		return
	}

	operations, err := read(instructions)
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(operations))
}

func read(n int) ([]string, error) {
	in := make([]string, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func solve(operations []string) int {
	var result int

	for _, op := range operations {
		if strings.Contains(op, "+") {
			result++
		} else if strings.Contains(op, "-") {
			result--
		}
	}

	return result
}

// Problem 282A: Bit++
// Time: 15m
