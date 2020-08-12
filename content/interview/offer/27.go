// +build OMIT

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

func prePrint(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, "-")
	prePrint(node.Left)
	prePrint(node.Right)
}

func main() {
	var root = &TreeNode{
		Val: 04,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
	prePrint(root)
	fmt.Println()
	m := mirrorTree(root)
	prePrint(m)
	fmt.Println()
}
