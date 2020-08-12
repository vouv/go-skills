// +build OMIT

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var result = make([]int, 0)
	recursion(head, &result)
	return result
}

func recursion(node *ListNode, result *[]int) {
	if node == nil {
		return
	}
	recursion(node.Next, result)
	*result = append(*result, node.Val)
}

func main() {
	var root = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(reversePrint(root))
}
