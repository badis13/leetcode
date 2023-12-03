package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var resTree *TreeNode

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}
	for {
		resTree.Val = root1.Val + root2.Val
		if root1.Left == nil && root2.Left == nil {
			if root1.Right == nil && root2.Right == nil {
				return resTree
			}
			root1 = root1.Right
			root2 = root2.Right
			resTree.Right = root1.Right
			resTree = resTree.Right
			continue
		}

		if root1.Right == nil && root2.Right == nil {
			root1 = root1.Left
			root2 = root2.Left
			resTree.Left = root1.Left
			resTree = resTree.Left
			continue
		}

		if root1.Left == nil {
			resTree.Left = root2.Left

		}

		if root2.Left == nil {
			resTree.Left = root1.Left
		}

		if root1.Right == nil {
			resTree.Right = root2.Right
		}

		if root2.Right == nil {
			resTree.Right = root1.Right
		}

	}

}

func main() {
	var head1 TreeNode
	var head2 TreeNode
	var first1 TreeNode
	var second1 TreeNode
	var third1 TreeNode
	var first2 TreeNode
	var second2 TreeNode
	var third2 TreeNode
	var fourth2 TreeNode
	head1.Val = 1
	first1.Val = 3
	second1.Val = 2
	third1.Val = 5
	head2.Val = 2
	first2.Val = 1
	second2.Val = 3
	third2.Val = 4
	fourth2.Val = 7
	head1.Left = &first1
	head1.Right = &second1
	first1.Left = &third1
	head2.Left = &first2
	head2.Right = &second2
	first2.Right = &third2
	second2.Right = &fourth2
	result := mergeTrees(&head1, &head2)
	fmt.Println(result)

}
