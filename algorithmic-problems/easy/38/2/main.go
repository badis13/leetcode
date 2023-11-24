package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	for p1, p2 := head.Next, head.Next.Next; ; p1, p2 = p1.Next, p2.Next.Next {
		if p1 == p2 {
			return true
		}
		if p1.Next == nil || p2.Next == nil || p2.Next.Next == nil {
			return false
		}
	}

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
	result := hasCycle(&head1)
	fmt.Println(result)

}
