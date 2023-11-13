package main

import (
	"fmt"
)

// heap solution https://leetcode.com/problems/top-k-frequent-elements/
func main() {
	nums := []int{-1, -1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	freqency := make(map[int]int)
	bucket := make([][]int, len(nums))
	result := make([]int, 0)

	if k == 0 {
		return result
	}

	if k == len(nums) {
		return nums
	}

	for _, value := range nums {
		freqency[value]++
	}
	if k == len(freqency) {
		for key, _ := range freqency {

			result = append(result, key)

		}
		return result
	}

	for key, val := range freqency {
		bucket[val] = append(bucket[val], key)
	}

	for i := len(bucket) - 1; i > 0 && len(result) < k; i-- {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
	}

	return result[:k]
}
