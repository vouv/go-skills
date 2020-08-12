// +build OMIT

package main

import "fmt"

func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _, num := range nums {
		a = (a ^ num) & (^b)
		b = (b ^ num) & (^a)
	}
	return a
}

func main() {
	fmt.Println("expect 1 and get", singleNumber(([]int{9, 1, 7, 9, 7, 9, 7})))
}
