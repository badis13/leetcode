package main

import "fmt"

// https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/

func main() {
	a := []int{1, 5, 2, 8, 2, 9}
	b := 9
	fmt.Println(maximizeSum(a, b))
}

func maximizeSum(nums []int, k int) int {
	var maxNum int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= nums[maxNum] {
			maxNum = i
		}
	}
	return ((2*nums[maxNum] + (k - 1)) * k) / 2

}
