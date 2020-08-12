package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	step := 0
	for len(stack) > 0 {
		var next []*TreeNode
		var arr []int
		for _, node := range stack {
			if step&1 == 0 {
				arr = append(arr, node.Val)
			} else {
				arr = append([]int{node.Val}, arr...)
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		step++
		stack = next
		result = append(result, arr)
	}
	return result
}

func main() {
	var root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	result := levelOrder(root)
	for _, level := range result {
		fmt.Println(level)
	}
}
