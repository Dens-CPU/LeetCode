package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//Если присутвует только 1 потомок
	if root.Left == nil {
		return 1 + MinDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + MinDepth(root.Left)
	}
	return 1 + min(MinDepth(root.Left), MinDepth(root.Right))
}

func BSF(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}

	return depth
}

func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 7}}}
	fmt.Println(MinDepth(&root))
	fmt.Println(BSF(&root))
}
