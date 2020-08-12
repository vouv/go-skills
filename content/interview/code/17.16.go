// +build OMIT

package main

import "fmt"

func massage(nums []int) int {
	var a, b int
	for _, v := range nums {
		a, b = b, max(b, a+v)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
}
