// +build OMIT

package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	cur := myPow(x, n>>1)
	return cur * cur * myPow(x, n&1)
}

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2.1, 3))
}
