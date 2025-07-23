package main

import (
	"fmt"
	"strconv"
)

func main() {
	var problems int
	var opnions []string

	fmt.Scan(&problems)
	if problems < 1 || problems > 1000 {
		fmt.Println("Invalid number of problems")
		return
	}

	opnions, err := read(problems * 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(clean(" ", opnions)))
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

func solve(opnions []int) int {
	var problemsToSolve int
	var sum int

	for i, op := range opnions {
		if op < 0 || op > 1 {
			panic("Invalid number on opnions. Must be 0 or 1")
		}
		sum += op
		if (i+1)%3 == 0 {
			if sum >= 2 {
				problemsToSolve++
				sum = 0
			} else {
				sum = 0
			}
		}
	}

	return problemsToSolve
}

func clean(char string, opnions []string) []int {
	var newOpnions []int
	for _, item := range opnions {
		if item != char {
			newOp, _ := strconv.Atoi(item)
			newOpnions = append(newOpnions, newOp)
		}
	}
	return newOpnions
}

// Problem 231A: Team
// Time: 1h
