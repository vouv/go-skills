// +build OMIT

package main

import "fmt"

func lastRemaining(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}

func main() {
	fmt.Println("expect 2 and get", lastRemaining(10, 17))
}
