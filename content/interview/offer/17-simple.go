// +build OMIT

package main

import "fmt"

var dict = [...]int{
	1: 1e1, 2: 1e2, 3: 1e3, 4: 1e4, 5: 1e5, 6: 1e6, 7: 1e7,
	9: 1e8, 10: 1e9, 11: 1e10, 12: 1e11, 13: 1e12, 14: 1e13,
}

func printNumbers(n int) []int {
	var result = make([]int, dict[n]-1)
	for i := range result {
		result[i] = i + 1
	}
	return result
}

func main() {
	fmt.Println(printNumbers(1))
}
