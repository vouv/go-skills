// +build OMIT

package main

import "fmt"

func lastRemaining(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

func main() {
	fmt.Println("expect 2 and get", lastRemaining(10, 17))
}
