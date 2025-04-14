package main

import "fmt"

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	for current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func main() {
	head := NewNode(1)
	current := head
	for i := 2; i <= 4; i++ {
		new := NewNode(2)
		current.Next = new
		current = current.Next
	}
	fmt.Println(head)
	current = head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	deleteDuplicates(head)
	fmt.Println()
	current = head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
}
