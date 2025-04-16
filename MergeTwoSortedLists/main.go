package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := &ListNode{}
	current3 := list3
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current3.Next = list1
			list1 = list1.Next
		} else {
			current3.Next = list2
			list2 = list2.Next
		}

		current3 = current3.Next
	}
	if list1 != nil {
		current3.Next = list1
	}
	if list2 != nil {

		current3.Next = list2
	}
	return list3.Next
}

func main() {
	var list1 ListNode
	var list2 ListNode
	list1 = ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list2 = ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergeTwoLists(&list1, &list2)
}
