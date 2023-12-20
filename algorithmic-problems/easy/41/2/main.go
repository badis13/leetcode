package main

import "fmt"

//https://leetcode.com/problems/average-of-levels-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type averageSum struct {
	sum  int
	nums int
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64

	if root.Left == nil && root.Right == nil {
		result = append(result, float64(root.Val))
		return result
	}

	resultMap := getAverage(root, 0)

	for i := 0; i < len(resultMap); i++ {
		result = append(result, float64(resultMap[i].sum)/float64(resultMap[i].nums))
	}

	return result
}

func getAverage(root *TreeNode, level int) map[int]averageSum {
	curMap := make(map[int]averageSum)
	curMap[level] = averageSum{curMap[level].sum + root.Val, curMap[level].nums + 1}

	if root.Left != nil {
		curMap = mergeMaps(curMap, getAverage(root.Left, level+1))
	}

	if root.Right != nil {
		curMap = mergeMaps(curMap, getAverage(root.Right, level+1))
	}

	return curMap
}

func mergeMaps(m1 map[int]averageSum, m2 map[int]averageSum) map[int]averageSum {
	merged := make(map[int]averageSum)
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		_, ok := merged[key]
		if ok {
			merged[key] = averageSum{merged[key].sum + value.sum, merged[key].nums + value.nums}
			continue
		}
		merged[key] = value
	}
	return merged
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
