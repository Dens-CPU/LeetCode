package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Convert(root *TreeNode) {
	if root == nil {
		fmt.Println("Пустое дерево")
		return
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("root=%d\n", root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if len(stack) == 0 {
			return
		}

		root.Right = stack[len(stack)-1]
		root.Left = nil
		fmt.Printf("root.Right=%d\n", root.Right.Val)
	}
}

func Recurs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return (root.Left)
	}
	if root.Right != nil {
		return (root.Right)
	}
	return root
}
func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	Recurs(&root)
	current := &root
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Right
	}
}
