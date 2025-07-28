package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		length := 0

		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			length++
		}

		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

// Problem 3: Longest Substring Without Repeating Characters
// Time: 30m
