package main

import "fmt"

//https://leetcode.com/problems/merge-sorted-array/description/

func main() {
	nums1 := []int{0, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 4, 5, 6}
	m := 0
	n := 6
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
			continue
		}

		nums1[m+n-1] = nums2[n-1]
		n--
	}

}
