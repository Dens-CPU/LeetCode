package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumElements(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + SumElements(root.Left) + SumElements(root.Right)
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 10}}
	fmt.Println(SumElements(root))
}
