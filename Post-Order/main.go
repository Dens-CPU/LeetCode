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
	var lastVisited *TreeNode

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && lastVisited != peek.Right {
				current = peek.Right
			} else {
				fmt.Printf("%d ", peek.Val)
				lastVisited = peek
				stack = stack[:len(stack)-1]
			}
		}

	}

}
func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}
	PostOrder(&root)
	fmt.Println()
	PostOrderIter(&root)
}
