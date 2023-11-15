package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/summary-ranges/

func main() {

	b := []int{0, 1, 2, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(summaryRanges(b))
}

func summaryRanges(nums []int) []string {
	result := make([]string, 0, len(nums))
	if len(nums) == 1 {
		result = append(result, strconv.Itoa(nums[0]))
		return result
	}
	cur := 0
	cur2 := 0
	for i := 0; i < len(nums)-1; i++ {
		if cur2 > 0 {
			result = append(result, strconv.Itoa(nums[i+1]))
			break
		}
		if nums[i] == nums[i+1]-1 {
			cur++
			if i+1 == len(nums)-1 {
				result = append(result, strconv.Itoa(nums[i+1-cur])+"->"+strconv.Itoa(nums[i+1]))
			}
			continue
		}

		if cur > 0 {
			result = append(result, strconv.Itoa(nums[i-cur])+"->"+strconv.Itoa(nums[i]))
			cur = 0
			if i+1 == len(nums)-1 {
				i--
				cur2++
			}
			continue
		}

		result = append(result, strconv.Itoa(nums[i]))
		cur = 0
		if i+1 == len(nums)-1 {
			i--
			cur2++
		}
	}
	return result
}
