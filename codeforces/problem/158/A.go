package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	scores := make([]int, n)
	for i := range scores {
		fmt.Scan(&scores[i])
	}

	passing := scores[k-1]
	count := 0

	for _, s := range scores {
		if s >= passing && s > 0 {
			count++
		}
	}

	fmt.Println(count)
}

// Problem 158A: Next Round
// Time: 25m
