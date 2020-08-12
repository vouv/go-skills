// +build OMIT

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var p = head
	for p != nil {
		next := p.Next

		node := &Node{
			Val:    p.Val,
			Next:   p.Next,
			Random: p.Random,
		}
		p.Next = node

		p = next
	}

	// 先处理random，不能动Next
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Next.Random.Next
		}
		p = p.Next.Next
	}

	// 还原head，分离copy的head
	ch := head.Next
	p = head
	cp := ch
	for p != nil {
		p.Next = p.Next.Next
		p = p.Next
		if cp.Next != nil {
			cp.Next = cp.Next.Next
			cp = cp.Next
		}
	}

	return ch
}

func main() {
	head := gen()
	cp := copyRandomList(head)
	for p, q := head, cp; p != nil && q != nil; p, q = p.Next, q.Next {
		fmt.Print(p.Val, p != q, p.Val == q.Val, " ")
		if p.Random != nil {
			fmt.Println(p.Random.Val == q.Random.Val)
		} else {
			fmt.Println(p.Random != q.Random)
		}
	}
}

func gen() *Node {
	node1 := &Node{
		Val: 7,
	}
	node2 := &Node{
		Val: 13,
	}
	node3 := &Node{
		Val: 11,
	}
	node4 := &Node{
		Val: 10,
	}
	node5 := &Node{
		Val: 1,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1
	return node1
}
