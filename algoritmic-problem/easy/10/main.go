package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func main() {
	a := []int{6, 4, 1, 65, 3}
	b := 5
	fmt.Println(twoSum(a, b))
}

func twoSum(nums []int, target int) []int {
	indices := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		indices[nums[i]] = append(indices[nums[i]], i)
	}

	for num, numIndices := range indices {
		pairNum := target - num
		if num == pairNum {
			if len(numIndices) > 1 {
				return numIndices[:2]
			}
			continue

		}

		pairIndices := indices[pairNum]
		if len(pairIndices) > 0 {
			return []int{numIndices[0], pairIndices[0]}
		}
	}

	return nil
}
