package main

import "fmt"

//https://leetcode.com/problems/average-of-levels-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64

	if root.Left == nil && root.Right == nil {
		result = append(result, float64(root.Val))
		return result
	}

	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		var nextLevel []*TreeNode
		sum := 0
		for _, item := range currentLevel {
			sum += item.Val
			if item.Left != nil {
				nextLevel = append(nextLevel, item.Left)
			}
			if item.Right != nil {
				nextLevel = append(nextLevel, item.Right)
			}
		}
		result = append(result, float64(sum)/float64(len(currentLevel)))
		currentLevel = nextLevel
	}

	return result
}

func main() {
	var head1 TreeNode
	var first1 TreeNode
	var second1 TreeNode
	var third1 TreeNode
	var fourth1 TreeNode
	var fifth1 TreeNode
	var sixth1 TreeNode
	head1.Val = 3
	first1.Val = 1
	second1.Val = 5
	third1.Val = 0
	fourth1.Val = 2
	fifth1.Val = 4
	sixth1.Val = 6
	head1.Left = &first1
	head1.Right = &second1
	first1.Left = &third1
	first1.Right = &fourth1
	second1.Left = &fifth1
	second1.Right = &sixth1
	result := averageOfLevels(&head1)
	fmt.Println(result)
}
