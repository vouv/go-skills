// +build OMIT

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA = 0
	for p := headA; p != nil; p = p.Next {
		lenA++
	}
	var lenB = 0
	for p := headB; p != nil; p = p.Next {
		lenB++
	}
	var p, q = headA, headB
	if lenA < lenB {
		lenA, lenB = lenB, lenA
		p, q = q, p
	}

	var diff = lenA - lenB
	for i := 0; i < diff; i++ {
		p = p.Next
	}

	for p != nil && q != nil && p != q {
		p = p.Next
		q = q.Next
	}
	return p
}

func main() {
	common := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: common,
		},
	}
	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: common,
			},
		},
	}
	fmt.Println("expect 8 and get", getIntersectionNode(headA, headB))
}
