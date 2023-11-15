package main

import "fmt"

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var res int
	for head != nil {
		res <<= 1
		res |= head.Val
		head = head.Next
	}

	return res
}

func main() {
	var head ListNode
	var first ListNode
	var second ListNode
	head.Val = 1
	first.Val = 0
	second.Val = 1
	head.Next = &first
	first.Next = &second

	fmt.Println(getDecimalValue(&head))
}
