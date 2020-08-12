// +build OMIT

package main

import "fmt"

func numWays(n int) int {
	const mod = 1e9 + 7
	f, g := 1, 1
	for i := 0; i < n-1; i++ {
		f, g = g, (f+g)%mod
	}
	return g
}

func main() {
	fmt.Println("expect numWays(95) = 93363621")
	fmt.Println("get numWays(95) =", numWays(95))
}
