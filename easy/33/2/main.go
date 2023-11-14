package main

import "fmt"

// https://leetcode.com/problems/single-number/

func main() {

	b := []int{1, 0, 1}
	fmt.Println(singleNumber(b))
}

func singleNumber(nums []int) int {
	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[j]^nums[i] == 0 {
			nums[j], nums[i], nums[len(nums)-1], nums[len(nums)-2] = nums[len(nums)-1], nums[len(nums)-2], nums[j], nums[i]
			nums = nums[:len(nums)-2]
			j = 0
			i = 0
			continue
		}

		if len(nums) > 2 {
			j++
			i = j
			continue
		}

	}
	return nums[0]
}
