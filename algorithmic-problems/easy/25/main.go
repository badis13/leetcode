package main

import "fmt"

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func curPostorder(root *Node, res []int) []int {
	for i := 0; i < len(root.Children); i++ {
		res = curPostorder(root.Children[i], res)
	}
	res = append(res, root.Val)
	return res
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 32)
	return curPostorder(root, res)
}

func main() {
	var head1 Node
	var child1 Node
	var child2 Node
	var child3 Node
	head1.Val = 5
	child1.Val = 2
	child2.Val = 8
	child3.Val = 2
	var head1Arr []*Node
	head1.Children = head1Arr
	fmt.Println(postorder(&head1))
}
