package main

import "fmt"

// https://leetcode.com/problems/single-number/

func main() {

	b := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(b))
}

func singleNumber(nums []int) int {
	curMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		curMap[nums[i]]++
	}
	resMap := make(map[int]int, len(curMap))
	for key, val := range curMap {
		resMap[val] = key
	}

	return resMap[1]

}
