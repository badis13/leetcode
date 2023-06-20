package main

import "fmt"

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var resultList ListNode
	var tail *ListNode = &resultList
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail, list1 = transferNode(tail, list1)
			continue
		}
		tail, list2 = transferNode(tail, list2)
	}

	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}

	return resultList.Next
}

func transferNode(destinationTail, sourceHead *ListNode) (*ListNode, *ListNode) {
	destinationTail.Next = sourceHead
	destinationTail = destinationTail.Next
	sourceHead = sourceHead.Next
	return destinationTail, sourceHead
}

func main() {
	var head1 ListNode
	var head2 ListNode
	var first1 ListNode
	var second1 ListNode
	var third1 ListNode
	var first2 ListNode
	var second2 ListNode
	var third2 ListNode
	head1.Val = 1
	first1.Val = 4
	second1.Val = 9
	third1.Val = 10
	head2.Val = 1
	first2.Val = 2
	second2.Val = 7
	third2.Val = 11
	head1.Next = &first1
	first1.Next = &second1
	second1.Next = &third1
	head2.Next = &first2
	first2.Next = &second2
	second2.Next = &third2
	result := mergeTwoLists(&head1, &head2)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}

}
