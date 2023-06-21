package main

import "fmt"

// https://leetcode.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var length int
	var current ListNode
	var tail *ListNode = &current
	for head != nil {
		tail, head = transferNode(tail, head)
		length++
	}
	length /= 2
	for i := 0; i <= length; i++ {
		current = *current.Next
	}

	return &current
}

func transferNode(destinationTail, sourceHead *ListNode) (*ListNode, *ListNode) {
	destinationTail.Next = sourceHead
	return destinationTail.Next, sourceHead.Next
}

func main() {
	var head1 ListNode
	var first1 ListNode
	var second1 ListNode
	var third1 ListNode
	head1.Val = 1
	first1.Val = 4
	second1.Val = 9
	third1.Val = 10
	head1.Next = &first1
	first1.Next = &second1
	second1.Next = &third1
	result := middleNode(&head1)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}

}
