// +build OMIT

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// 前序序列第一个元素为根节点
	root := preorder[0]
	// 寻找根节点在中序序列中的位置
	// 从而找出左右子树的前序序列和中序序列
	idx := 0
	for i, v := range inorder {
		if v == root {
			idx = i
			break
		}
	}
	node := &TreeNode{}
	node.Val = root

	// 递归寻找和构建左右子树
	node.Left = buildTree(preorder[1:1+idx], inorder[0:idx])
	node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return node
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	root := buildTree(pre, in)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}
