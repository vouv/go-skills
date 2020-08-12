// +build OMIT

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: nil,
	}
	for p := head; p != nil; {
		n := p.Next
		p.Next = dummy.Next
		dummy.Next = p
		p = n
	}
	return dummy.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	show(head)

	res := reverseList(head)

	show(res)
}

func show(node *ListNode) {
	for p := node; p != nil; p = p.Next {
		if p.Next == nil {
			fmt.Println(p.Val)
		} else {
			fmt.Print(p.Val, "-")
		}
	}
}
