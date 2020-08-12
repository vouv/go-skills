// +build OMIT

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var head, pre *Node

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	dfs(root)
	head.Left = pre
	pre.Right = head
	return head
}

func dfs(node *Node) {
	if node == nil {
		return
	}
	dfs(node.Left)
	if pre == nil {
		head = node
	} else {
		pre.Right = node
		node.Left = pre
	}
	pre = node
	dfs(node.Right)
	return
}

func main() {
	var tree = &Node{
		Val: 4,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 1},
			Right: &Node{Val: 3},
		},
		Right: &Node{
			Val: 5,
		},
	}
	list := treeToDoublyList(tree)

	var i int
	for p := list; i < 10; p = p.Right {
		fmt.Println(p.Val)
		i++
	}
	i = 0
	for p := list; i < 10; p = p.Left {
		fmt.Println(p.Val)
		i++
	}
}
