// +build OMIT

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p, q := head, head
	for i := 1; i < k; i++ {
		if q.Next != nil {
			q = q.Next
		} else {
			return nil
		}
	}
	for q.Next != nil {
		q = q.Next
		p = p.Next
	}
	return p
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

	fmt.Println("before:")
	show(head)

	res := getKthFromEnd(head, 3)

	fmt.Println("after:")
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
