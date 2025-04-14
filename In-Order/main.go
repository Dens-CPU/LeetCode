package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Рекурсия
func InorderTraversal(root *TreeNode) []int {
	var array []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		array = append(array, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return array
}

// Итерации
func InorderTraversalIter(root *TreeNode) []int {
	array := []int{}
	stack := []*TreeNode{}
	current := root
	for current != nil || len(stack) > 0 {
		//Продвигаемя максимально глубоко по левому поддереву
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		array = append(array, current.Val)
		//Переход на правое поодерево
		current = current.Right
	}
	return array
}
func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}
	fmt.Println(InorderTraversal(&root))
	fmt.Println(InorderTraversalIter(&root))
}
