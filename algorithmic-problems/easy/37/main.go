package main

import "fmt"

//https://leetcode.com/problems/concatenation-of-array/submissions/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
