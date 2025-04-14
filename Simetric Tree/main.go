package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Итерационный метод
// func isSymmetric(root *TreeNode) bool {
// 	stack := []*TreeNode{}
// 	stack = append(stack, root.Left, root.Right)
// 	for len(stack) > 0 {
// 		l, r := stack[0], stack[1]
// 		stack = stack[2:]
// 		if l == nil && r == nil {
// 			return true
// 		}
// 		if (l != nil && r == nil) || (l == nil && r != nil) {
// 			return false
// 		}
// 		if l.Val != r.Val {
// 			return false
// 		}
// 		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
// 	}
// 	return true
// }

// Рекурсия
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}
func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

// Очередь
func BSF(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := [][2]*TreeNode{{root.Left, root.Right}}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		left, right := pair[0], pair[1]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, [2]*TreeNode{left.Left, right.Right})
		queue = append(queue, [2]*TreeNode{left.Right, right.Left})

	}
	return true
}
func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	fmt.Println(isSymmetric(&root))
	fmt.Println(BSF(&root))
}
