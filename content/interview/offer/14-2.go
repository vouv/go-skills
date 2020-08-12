// +build OMIT

package main

import "fmt"

func cuttingRope(n int) int {
	const mod = 1e9 + 7
	var tmp = []int{0, 0, 1, 2, 4, 6, 9}
	// 对于大数只需要拆分到2或3
	if n < 7 {
		return tmp[n]
	}
	return cuttingRope(n-3) * 3 % mod
}

func main() {
	fmt.Println("expect 729 and get", cuttingRope(18))
}
