// +build OMIT

package main

import "fmt"

func findNthDigit(n int) int {
	// 当前数的位数
	digit := 1
	// 当前数的开始
	start := 1
	// 9 * 1、 90 * 2、 900 * 3...
	count := 9
	// 求出所在区间
	for n > count {
		n -= count
		start *= 10
		digit++
		count = digit * 9 * start
	}
	// 目标数，索引从0开始所以要n-1
	target := start + (n-1)/digit
	// 在目标串中的位置，索引从0开始所以要n-1
	idx := (n - 1) % digit
	return int(strconv.Itoa(target)[idx] - '0')
}

func main() {
	fmt.Println("expect 5 and get", findNthDigit(7896))
}
