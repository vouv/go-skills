// +build OMIT

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) Print() {
	iter := n
	for iter != nil {
		if iter.Next == nil {
			fmt.Print(iter.Val, "\n")
		} else {
			fmt.Print(iter.Val, "-")
		}
		iter = iter.Next
	}
}

func Reverse(head *Node) *Node {
	var result = &Node{}
	for p := head; p != nil; {
		next := p.Next
		p.Next = result.Next
		result.Next = p
		p = next
	}
	return result.Next
}

func main() {
	head := &Node{
		Val: 0,
		Next: &Node{
			Val: 1,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 4,
						Next: &Node{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Print("list:    ")
	head.Print()
	fmt.Print("reverse: ")
	Reverse(head).Print()
}
