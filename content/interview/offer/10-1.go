// +build OMIT

package main

import "fmt"

func fib(n int) int {
	const mod = 1e9 + 7
	f, g := 0, 1
	for i := 0; i < n; i++ {
		f, g = g, (f+g)%mod
	}
	return f
}

func main() {
	fmt.Println("expect fib(95) = 407059028")
	fmt.Println("get fib(95) =", fib(95))
}
