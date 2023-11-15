package main

import "fmt"

// https://leetcode.com/problems/missing-number/
func main() {
	b := []int{2, 0, 1, 5, 3}
	fmt.Println(missingNumber(b))
}

func missingNumber(nums []int) (result int) {
	for i := 0; i < len(nums); i++ {
		result -= nums[i] - i - 1
	}

	return result
}
