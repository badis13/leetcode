package main

import "fmt"

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func curPreorder(root *Node, res []int) []int {
	res = append(res, root.Val)
	for i := 0; i < len(root.Children); i++ {
		res = curPreorder(root.Children[i], res)
	}
	return res
}

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return nil
	}
	return curPreorder(root, res)
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
	fmt.Println(preorder(&head1))
}
