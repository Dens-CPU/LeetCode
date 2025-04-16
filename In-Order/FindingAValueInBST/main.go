package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindValue(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.Val {
		return true
	}
	if root.Left != nil && val < root.Val {
		return FindValue(root.Left, val)
	}
	if root.Right != nil && val > root.Val {

		return FindValue(root.Right, val)
	}
	return false
}

func main() {
	root := TreeNode{Val: 10, Left: &TreeNode{Val: 16, Left: &TreeNode{Val: 21}, Right: &TreeNode{Val: 13}}, Right: &TreeNode{Val: 5}}
	val := 55
	fmt.Println(FindValue(&root, val))
}
