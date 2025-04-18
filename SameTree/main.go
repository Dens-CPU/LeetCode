package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//DFS - обход дерева в глубину
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
	q := TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
	fmt.Println(isSameTree(&p, &q))

}
