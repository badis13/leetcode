package main

import "fmt"

// https://leetcode.com/problems/single-number/

func main() {

	b := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(b))
}

func singleNumber(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		result ^= nums[i]

	}
	return result
}
