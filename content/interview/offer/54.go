// +build OMIT

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	var result int
	recursion(root, &k, &result)
	return result
}

func recursion(node *TreeNode, k, result *int) {
	if node == nil {
		return
	}
	recursion(node.Right, k, result)
	*k--
	if *k == 0 {
		*result = node.Val
		return
	}
	recursion(node.Left, k, result)
}

func main() {

}
