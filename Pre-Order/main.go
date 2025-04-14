package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Рекурсия
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

//Итерация
func PreOrderIter(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		fmt.Printf("%d ", root.Val)
		stack = stack[:len(stack)-1]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
}
func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}
	PreOrder(&root)
	fmt.Println()
	PreOrderIter(&root)

}
