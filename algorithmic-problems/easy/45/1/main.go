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
	if n == 0 {
		return
	}

	tailRes := m + n - 1
	tail1 := m - 1
	tail2 := n - 1

	if tail1 < 0 {
		copy(nums1, nums2)
		return
	}

	for tail1 >= 0 && tail2 >= 0 {
		if nums1[tail1] > nums2[tail2] {
			nums1[tailRes] = nums1[tail1]
			if tailRes == 1 {
				if nums1[tail1] > nums2[tail2] {
					nums1[0] = nums2[0]
					return
				}
				nums1[tailRes] = nums2[tail2]
				return
			}
			tailRes--
			tail1--
			for tailRes >= 0 && tail1 < 0 {
				nums1[tailRes] = nums2[tail2]
				tail2--
				tailRes--
			}
			continue
		}
		nums1[tailRes] = nums2[tail2]
		if tailRes == 1 {
			if nums1[tail1] > nums2[tail2] {
				return
			}
			nums1[tailRes] = nums2[tail2]
			return
		}
		tail2--
		tailRes--

	}

}
