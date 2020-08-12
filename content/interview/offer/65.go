// +build OMIT

package main

import "fmt"

func add(a int, b int) int {
	for b != 0 {
		// a ^ b 表示a与b的无进位相加
		// (a & b) << 1 表示求a与b相加的进位，与之后需要左移到进位上
		a, b = a^b, (a&b)<<1
	}
	return a
}

func main() {
	fmt.Println(add(66, 22))
}
