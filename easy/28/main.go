package main

import "fmt"

// https://leetcode.com/problems/remove-element/

func main() {
	nums := []int{1, 1, 2, 3, 4, 4, 4, 5}
	k := 4
	fmt.Println(removeElement(nums, k))
}

func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
			continue
		}
	}
	return result
}
