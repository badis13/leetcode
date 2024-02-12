package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) (result int) {
	if len(nums) == 1 {
		return nums[0]
	}
	majority := len(nums) / 2
	curRes := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		curRes[nums[i]]++
	}

	for key, val := range curRes {
		if (val > result) && (val > majority) {
			result = key
		}
	}

	return result
}
