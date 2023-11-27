package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func main() {
	a := []int{8, 1, 2, 2, 3}

	fmt.Println(smallerNumbersThanCurrent(a))
}

func smallerNumbersThanCurrent(nums []int) []int {
	cur := make([]int, len(nums))
	copy(cur, nums)
	result := make([]int, len(nums))
	sort.Ints(cur)
	sortCur := make(map[int]int)

	for index, val := range cur {
		_, ok := sortCur[val]
		if !ok {
			sortCur[val] = index
		}
	}

	for index, val := range nums {
		result[index] = sortCur[val]
	}

	return result

}
