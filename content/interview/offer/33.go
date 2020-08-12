// +build OMIT

package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n == 0 {
		return true
	}
	root := postorder[n-1]

	// 从p往右（包含p）是右子树
	var p = 0
	for p < n-1 && postorder[p] < root {
		p++
	}

	for i := p; i < n-1; i++ {
		if postorder[i] < root {
			return false
		}
	}

	var lok = true
	if p > 0 {
		lok = verifyPostorder(postorder[0:p])
	}

	var rok = true
	if p < n-1 {
		rok = verifyPostorder(postorder[p : n-1])
	}

	return lok && rok
}

func main() {
	fmt.Println("expect false and get", verifyPostorder([]int{1, 6, 3, 2, 5}))
	fmt.Println("expect true and get", verifyPostorder([]int{1, 3, 2, 6, 5}))
}
