package main

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

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

}
