package main

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

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

}
