package main

import "fmt"

// https://leetcode.com/problems/find-the-difference-of-two-arrays/

func main() {
	a := []int{1, 1, 2, 3, 4}
	b := []int{11, 23, 2, 4}
	fmt.Println(findDifference(a, b))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	result := [][]int{}
	setNums1 := make(map[int]bool)
	setNums2 := make(map[int]bool)
	resArr1 := []int{}
	resArr2 := []int{}
	for _, value := range nums1 {
		setNums1[value] = true
	}
	for _, value := range nums2 {
		setNums2[value] = true
	}
	for key := range setNums1 {
		if !setNums2[key] {
			resArr1 = append(resArr1, key)
		}
	}
	for key := range setNums2 {
		if !setNums1[key] {
			resArr2 = append(resArr2, key)
		}
	}

	result = append(result, resArr1, resArr2)
	return result
}
