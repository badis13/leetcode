package main

import "fmt"

//https://leetcode.com/problems/contains-duplicate/description/

func main() {
	nums := []int{1, 2, 3, 1, 11, 12, 19, 13, 29, 33, 21, 24}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {

	if len(nums) < 2 {
		return false
	}

	hashSort := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hashSort[nums[i]]++
	}

	return len(hashSort) < len(nums)
}
