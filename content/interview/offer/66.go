// +build OMIT

package main

import "fmt"

func constructArr(a []int) []int {
	n := len(a)
	var result = make([]int, n)
	if n == 0 {
		return result
	}
	// 计算a[0] * ... * a[i]
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * a[i-1]
	}

	// 计算a[i+1] * ... * a[n-1]
	var tmp = 1
	for i := n - 2; i >= 0; i-- {
		tmp *= a[i+1]
		result[i] *= tmp
	}
	return result
}

func main() {
	fmt.Println("expect [120,60,40,30,24] and get",
		constructArr([]int{1, 2, 3, 4, 5}))
}
