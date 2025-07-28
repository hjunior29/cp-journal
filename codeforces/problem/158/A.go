package main

import "fmt"

func main() {
	var competitors []int
	var scores []int

	competitors, err := read(2)
	if err != nil {
		panic(err)
	}

	n := competitors[0]
	k := competitors[1]

	scores, err = read(n)
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(n, k, scores))
}

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func solve(n, k int, scores []int) int {
	passingGrade := scores[k-1]
	var passed int

	for _, s := range scores {
		if s >= passingGrade && s > 0 {
			passed++
		}
	}

	return passed
}

// Problem 158A: Next Round
// Time: 25m
