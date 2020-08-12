// +build OMIT

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}
	return match(a, b) || isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func match(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a.Val == b.Val && match(a.Left, b.Left) && match(a.Right, b.Right)
}

func main() {
	var a = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	var b = &TreeNode{
		Val:  4,
		Left: &TreeNode{Val: 1},
	}
	fmt.Println(isSubStructure(a, b))
}
