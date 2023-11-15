package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	result := 0
	for first := 0; first < len(nums); first++ {
		if nums[first] != nums[result] {
			result++
			nums[result] = nums[first]
		}
	}
	return result + 1
}
