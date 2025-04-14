package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d ", root.Val)
}

func PostOrderIter(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{}
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			if current.Right != nil {
				stack = append(stack, current, current.Right)
			} else {
				stack = append(stack, current)
			}
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", current.Val)
		current = current.Right
	}
}
func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}
	PostOrder(&root)
	fmt.Println()
	PostOrderIter(&root)
}
