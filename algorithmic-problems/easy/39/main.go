package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func main() {
	a := []int{1, 2, 2, 5, 12, 7}

	fmt.Println(smallerNumbersThanCurrent(a))
}

func smallerNumbersThanCurrent(nums []int) []int {
	cur := make([]int, len(nums))
	copy(cur, nums)
	result := make([]int, len(nums))
	sort.Ints(cur)
	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums) {
			result[i] = i
			return result
		}

		if i > 0 && cur[i] == cur[i-1] {
			result[i] = result[i-1]
			continue
		}

		if cur[i] == cur[i+1] {
			result[i], result[i+1] = i, i
			continue
		}
		result[i] = i
	}

	return result

}
