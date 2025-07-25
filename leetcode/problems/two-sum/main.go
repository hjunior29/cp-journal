package main

import "fmt"

func main() {
	var nums []int
	var target int

	nums = []int{2, 7, 11, 15}
	target = 9

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	var sum int
	var twoSum []int

	for i, n1 := range nums {
		for j, n2 := range nums {
			if i == j {
				continue
			}
			sum = n1 + n2
			if sum == target {
				twoSum = []int{i, j}
				return twoSum
			}
		}
	}

	return nil
}

// Problem: Two Sum
// Time: 20m
