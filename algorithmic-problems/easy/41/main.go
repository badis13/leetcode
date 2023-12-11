package main

import "fmt"

//https://leetcode.com/problems/average-of-levels-in-binary-tree/
//не работает, нужно починить!!!

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	var left, right float64
	if root.Left == nil && root.Right == nil {
		return result
	}

	if root.Left != nil && root.Right != nil {
		left = float64(root.Left.Val)
		right = float64(root.Right.Val)
		result = append(result, (left+right)/2)
		result = append(result, averageOfLevels(root.Left)...)
		result = append(result, averageOfLevels(root.Right)...)
	}

	if root.Left != nil && root.Right == nil {
		result = append(result, float64(root.Left.Val))
		result = append(result, averageOfLevels(root.Left)...)
	}

	if root.Left == nil && root.Right != nil {
		result = append(result, float64(root.Right.Val))
		result = append(result, averageOfLevels(root.Right)...)
	}

	return result
}

func main() {
	var head1 TreeNode
	var first1 TreeNode
	var second1 TreeNode
	var third1 TreeNode
	var fourth1 TreeNode
	head1.Val = 3
	first1.Val = 9
	second1.Val = 20
	third1.Val = 15
	fourth1.Val = 7
	head1.Left = &first1
	head1.Right = &second1
	first1.Left = nil
	first1.Right = nil
	second1.Left = &third1
	second1.Left = &fourth1
	result := averageOfLevels(&head1)
	fmt.Println(result)
}
