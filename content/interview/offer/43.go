// +build OMIT

package main

import (
	"fmt"
)

func countDigitOne(n int) int {
	digit, res := 1, 0
	low, high := 0, n/10
	cur := n % 10

	for high != 0 || cur != 0 {
		// 高位 * 低位为1的个数
		// 然后通过对当前位进行判断，找真实的为1的个数
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		// 看低位有多少种为1的方式
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}

func main() {
	fmt.Println("expect 5 ang get", countDigitOne(12))
}
