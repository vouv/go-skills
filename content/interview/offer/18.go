// +build OMIT

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{
		Next: head,
	}
	// 双指针
	var p, q = dummy, head
	for q != nil && q.Val != val {
		q = q.Next
		p = p.Next
	}
	if p == nil {
		return dummy.Next
	}

	p.Next = q.Next
	return dummy.Next
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	fmt.Println("before:")
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	head = deleteNode(head, 5)

	fmt.Println("after:")
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
