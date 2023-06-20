package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev

}

func main() {
	var head ListNode
	var first ListNode
	var second ListNode
	head.Val = 1
	first.Val = 0
	second.Val = 5
	head.Next = &first
	first.Next = &second

	fmt.Println(reverseList(&head).Val)
}
