// +build OMIT

package main

import "fmt"

func sumNums(n int) int {
	var result = n
	helper(&result, n-1)
	return result
}

func helper(n *int, deep int) bool {
	*n += deep
	return deep > 0 && helper(n, deep-1)
}

func main() {
	fmt.Println("expect 5050 and get", sumNums(100))
}
