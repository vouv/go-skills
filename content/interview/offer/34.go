// +build OMIT

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result = [][]int{}
	recursion(root, sum, []int{}, &result)
	return result
}

func recursion(node *TreeNode, last int, buf []int, result *[][]int) {
	if node == nil {
		return
	}
	buf = append(buf, node.Val)
	if node.Left == nil && node.Right == nil && last == node.Val {
		res := make([]int, len(buf))
		copy(res, buf)
		*result = append(*result, res)
		return
	}
	recursion(node.Left, last-node.Val, buf, result)
	recursion(node.Right, last-node.Val, buf, result)
	buf = buf[:len(buf)-1]
}

func main() {
	var root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}
